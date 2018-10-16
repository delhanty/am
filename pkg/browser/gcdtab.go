package browser

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/linkai-io/am/am"
	"github.com/linkai-io/am/pkg/convert"
	"github.com/wirepair/gcd"
	"github.com/wirepair/gcd/gcdapi"
)

var (
	ErrNavigationTimedOut = errors.New("navigation timed out")
	ErrTabCrashed         = errors.New("tab crashed")
	ErrTabClosing         = errors.New("closing")
	ErrTimedOut           = errors.New("request timed out")
)

type Tab struct {
	t                     *gcd.ChromeTarget
	address               *am.ScanGroupAddress
	container             *ResponseContainer
	crashedCh             chan string
	exitCh                chan struct{}
	navigationCh          chan int
	lastNodeChangeTimeVal atomic.Value
}

func NewTab(tab *gcd.ChromeTarget, address *am.ScanGroupAddress) *Tab {
	t := &Tab{
		t:            tab,
		address:      address,
		container:    NewResponseContainer(),
		crashedCh:    make(chan string),
		exitCh:       make(chan struct{}),
		navigationCh: make(chan int),
	}
	t.subscribeBrowserEvents()
	return t
}

// LoadPage capture network traffic and take screen shot of DOM and image
func (t *Tab) LoadPage(ctx context.Context, url string) error {
	navParams := &gcdapi.PageNavigateParams{Url: url, TransitionType: "typed"}
	log.Info().Str("url", url).Msg("navigating")
	_, _, _, err := t.t.Page.NavigateWithParams(navParams)
	if err != nil {
		log.Warn().Err(err).Str("host_address", t.address.HostAddress).
			Str("ip_address", t.address.IPAddress).
			Str("url", url).Msg("failed to load page")
		return err
	}
	log.Info().Str("url", url).Msg("navigating complete")
	return t.WaitReady(ctx, time.Second*3)
}

// TakeScreenshot returns a png image, base64 encoded, or error if failed
func (t *Tab) TakeScreenshot(ctx context.Context) (string, error) {
	_, _, rect, err := t.t.Page.GetLayoutMetrics()
	if err != nil {
		return "", err
	}

	params := &gcdapi.PageCaptureScreenshotParams{
		Format:  "png",
		Quality: 100,
		Clip: &gcdapi.PageViewport{
			X:      rect.X,
			Y:      rect.Y,
			Width:  rect.Width,
			Height: rect.Height,
			Scale:  float64(1)},
		FromSurface: true,
	}

	return t.t.Page.CaptureScreenshotWithParams(params)
}

// Close the exit channel
func (t *Tab) Close() {
	close(t.exitCh)
}

// WaitReady waits for the page to load, DOM to be stable, and no network traffic in progress
func (t *Tab) WaitReady(ctx context.Context, stableAfter time.Duration) error {
	ticker := time.NewTicker(150 * time.Millisecond)
	defer ticker.Stop()

	navTimer := time.After(30 * time.Second)
	log.Info().Msg("waiting for nav to complete")
	// wait navigation to complete.
	select {
	case <-navTimer:
		return ErrNavigationTimedOut
	case <-ctx.Done():
		return ctx.Err()
	case <-t.exitCh:
		return errors.New("exiting")
	case reason := <-t.crashedCh:
		return errors.Wrap(ErrTabCrashed, reason)
	case <-t.navigationCh:
	}

	stableTimer := time.After(5 * time.Second)
	// wait for DOM & network stability
	log.Info().Msg("waiting for DOM & network stability")
	for {
		select {
		case reason := <-t.crashedCh:
			return errors.Wrap(ErrTabCrashed, reason)
		case <-ctx.Done():
			return ctx.Err()
		case <-t.exitCh:
			return ErrTabClosing
		case <-stableTimer:
			return ErrTimedOut
		case <-ticker.C:
			log.Info().Int32("requests", t.container.GetRequests()).Msg("tick")
			if changeTime, ok := t.lastNodeChangeTimeVal.Load().(time.Time); ok {
				if time.Now().Sub(changeTime) >= stableAfter && t.container.GetRequests() == 0 {
					// times up, should be stable now
					return nil
				}
			}
		}
	}
}

// SerializeDOM and return it as string
func (t *Tab) SerializeDOM() string {
	node, err := t.t.DOM.GetDocument(-1, true)
	if err != nil {
		return ""
	}
	html, err := t.t.DOM.GetOuterHTMLWithParams(&gcdapi.DOMGetOuterHTMLParams{
		NodeId: node.NodeId,
	})
	if err != nil {
		return ""
	}
	return html
}

// GetNetworkTraffic returns all responses after page load
func (t *Tab) GetNetworkTraffic() []*am.HTTPResponse {
	return t.container.GetResponses()
}

// InjectIP replaces the address.HostAddress with the IP address so we can catalogue all variants
// of the host/ip pairs.
// TODO: handle punycode when doing strings.Replace.
func (t *Tab) InjectIP(scheme, port string) {

	httpPattern := &gcdapi.NetworkRequestPattern{
		UrlPattern: "http://" + t.address.HostAddress + "*",
	}
	httpsPattern := &gcdapi.NetworkRequestPattern{
		UrlPattern: "https://" + t.address.HostAddress + "*",
	}

	patterns := make([]*gcdapi.NetworkRequestPattern, 2)
	patterns[0] = httpPattern
	patterns[1] = httpsPattern

	interceptParams := &gcdapi.NetworkSetRequestInterceptionParams{
		Patterns: patterns,
	}

	t.t.Network.SetRequestInterceptionWithParams(interceptParams)

	t.t.Subscribe("Network.requestIntercepted", func(target *gcd.ChromeTarget, payload []byte) {
		r := &gcdapi.NetworkRequestInterceptedEvent{}
		if err := json.Unmarshal(payload, r); err != nil {
			log.Warn().Err(err).Msg("failed to unmarshal network request intercepted")
			return
		}

		headers := r.Params.Request.Headers
		parsedURL, err := url.Parse(r.Params.Request.Url)
		if err != nil {
			headers["host"] = t.address.HostAddress + ":" + port
		} else {
			headers["host"] = parsedURL.Host // will return host:port
		}

		ipURL := strings.Replace(r.Params.Request.Url, t.address.HostAddress, t.address.IPAddress, 1)
		log.Info().Str("host_address", t.address.HostAddress).
			Str("ip_address", t.address.IPAddress).
			Msg("intercepted and replacing IP")
		p := &gcdapi.NetworkContinueInterceptedRequestParams{
			InterceptionId: r.Params.InterceptionId,
			Url:            ipURL,
			Headers:        headers,
		}
		target.Network.ContinueInterceptedRequestWithParams(p)
		log.Info().Str("host_address", t.address.HostAddress).
			Str("ip_address", t.address.IPAddress).
			Msg("continue called")
	})
}

// CaptureNetworkTraffic ensures we capture all traffic (only saving text bodies) during navigation.
func (t *Tab) CaptureNetworkTraffic(ctx context.Context, address *am.ScanGroupAddress, port string) {

	t.t.Network.EnableWithParams(&gcdapi.NetworkEnableParams{
		MaxPostDataSize:       -1,
		MaxResourceBufferSize: -1,
		MaxTotalBufferSize:    -1,
	})

	t.t.Subscribe("Network.requestWillBeSent", func(target *gcd.ChromeTarget, payload []byte) {

	})

	t.t.Subscribe("Network.responseReceived", func(target *gcd.ChromeTarget, payload []byte) {

		defer t.container.DecRequest()
		t.container.IncRequest()

		message := &gcdapi.NetworkResponseReceivedEvent{}
		if err := json.Unmarshal(payload, message); err != nil {
			return
		}

		timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		p := message.Params

		log.Info().Str("request_id", p.RequestId).Str("url", p.Response.Url).Msg("waiting")
		if err := t.container.WaitFor(timeoutCtx, p.RequestId); err != nil {
			return
		}

		response := t.buildResponse(address, port, message)
		t.container.Add(response)
	})

	t.t.Subscribe("Network.loadingFinished", func(target *gcd.ChromeTarget, payload []byte) {
		message := &gcdapi.NetworkLoadingFinishedEvent{}
		if err := json.Unmarshal(payload, message); err != nil {
			return
		}
		log.Info().Str("request_id", message.Params.RequestId).Msg("finished")
		t.container.BodyReady(message.Params.RequestId)
	})
}

// buildResponse fills out a new am.HTTPResponse with all relevant details
func (t *Tab) buildResponse(address *am.ScanGroupAddress, requestedPort string, message *gcdapi.NetworkResponseReceivedEvent) *am.HTTPResponse {
	var host string
	var responsePort string
	var scheme string

	p := message.Params
	u, err := url.Parse(p.Response.Url)
	if err != nil {
		log.Warn().
			Err(err).
			Str("host_address", address.HostAddress).
			Str("ip_address", address.IPAddress).
			Str("port", requestedPort).
			Msg("failed to parse url, results may be inaccurate")

		host = address.HostAddress
		responsePort = requestedPort
		scheme = "http"
	} else {
		host = u.Host
		responsePort = u.Port()
		scheme = u.Scheme
	}

	response := &am.HTTPResponse{
		Scheme:        scheme,
		RequestedPort: requestedPort,
		Host:          host,
		ResponsePort:  responsePort,
		RequestID:     p.RequestId,
		URL:           p.Response.Url,
		Headers:       p.Response.Headers,
		MimeType:      p.Response.MimeType,
		Status:        p.Response.Status,
		StatusText:    p.Response.StatusText,
		RawBody:       t.encodeResponseBody(message),
	}

	if p.Type == "Document" {
		response.IsDocument = true
		response.WebCertificate = t.extractCertificate(message)
	}

	response.ResponseTimestamp = time.Now().UnixNano()

	return response
}

func (t *Tab) extractCertificate(message *gcdapi.NetworkResponseReceivedEvent) *am.WebCertificate {
	p := message.Params

	u, err := url.Parse(p.Response.Url)
	if err != nil {
		return nil
	}

	if u.Hostname() == t.address.HostAddress && u.Scheme == "https" &&
		strings.HasPrefix(p.Response.Url, "https") && p.Response.SecurityDetails != nil {

		return convert.NetworkCertificateToWebCertificate(p.Response.SecurityDetails)
	}

	return nil
}

func (t *Tab) encodeResponseBody(p *gcdapi.NetworkResponseReceivedEvent) string {

	var err error
	var encoded bool
	var body []byte
	var bodyStr string

	bodyStr, encoded, err = t.t.Network.GetResponseBody(p.Params.RequestId)
	if err != nil {
		log.Warn().Str("url", p.Params.Response.Url).Err(err).Msg("failed to get body")
	}

	body = []byte(bodyStr)
	if encoded {
		body, _ = base64.StdEncoding.DecodeString(bodyStr)
	}

	// we don't want to capture anything other than text based files.
	if !strings.HasPrefix(http.DetectContentType(body), "text") {
		bodyStr = ""
	}

	return bodyStr
}

func (t *Tab) domUpdated() func(target *gcd.ChromeTarget, payload []byte) {
	return func(target *gcd.ChromeTarget, payload []byte) {
		log.Info().Msg("dom updated")
		t.lastNodeChangeTimeVal.Store(time.Now())
	}
}

func (t *Tab) subscribeBrowserEvents() {
	t.t.DOM.Enable()
	t.t.Inspector.Enable()
	t.t.Page.Enable()

	t.t.Subscribe("Inspector.targetCrashed", func(target *gcd.ChromeTarget, payload []byte) {
		select {
		case t.crashedCh <- "crashed":
		case <-t.exitCh:
		}
	})

	t.t.Subscribe("Inspector.detached", func(target *gcd.ChromeTarget, payload []byte) {
		header := &gcdapi.InspectorDetachedEvent{}
		err := json.Unmarshal(payload, header)
		reason := "detached"

		if err == nil {
			reason = header.Params.Reason
		}

		select {
		case t.crashedCh <- reason:
		case <-t.exitCh:
		}
	})

	t.t.Subscribe("Page.loadEventFired", func(target *gcd.ChromeTarget, payload []byte) {
		select {
		case t.navigationCh <- 0:
		case <-t.exitCh:
		}
	})

	// new nodes
	t.t.Subscribe("DOM.setChildNodes", t.domUpdated())
	t.t.Subscribe("DOM.attributeModified", t.domUpdated())
	t.t.Subscribe("DOM.attributeRemoved", t.domUpdated())
	t.t.Subscribe("DOM.characterDataModified", t.domUpdated())
	t.t.Subscribe("DOM.childNodeCountUpdated", t.domUpdated())
	t.t.Subscribe("DOM.childNodeInserted", t.domUpdated())
	t.t.Subscribe("DOM.childNodeRemoved", t.domUpdated())
	t.t.Subscribe("DOM.documentUpdated", t.domUpdated())

}
