package ns

import (
	"context"
	"errors"
	"runtime"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/linkai-io/am/pkg/cache"
	"github.com/linkai-io/am/pkg/parsers"
	"github.com/miekg/dns"

	"github.com/linkai-io/am/am"
	"github.com/linkai-io/am/pkg/convert"
	"github.com/linkai-io/am/pkg/dnsclient"
	"github.com/linkai-io/am/services/module"
	"github.com/linkai-io/am/services/module/ns/state"
)

const (
	// how long until we remove the zone record saying we've already done ns lookups
	nsExpire = 14400
)

var (
	// ErrEmptyDNSServer missing dns server
	ErrEmptyDNSServer = errors.New("dns_server was empty or invalid")
)

// NS module for extracting NS related information for a scan group.
type NS struct {
	st state.Stater
	dc *dnsclient.Client
	// for closing subscriptions to listen for group updates
	exitContext context.Context
	cancel      context.CancelFunc
	// concurrent safe cache of scan groups updated via Subscribe callbacks
	groupCache *cache.ScanGroupSubscriber
}

// New creates a new NS module for identifying zone information via DNS
// and storing the results in Redis.
func New(dc *dnsclient.Client, st state.Stater) *NS {
	ctx, cancel := context.WithCancel(context.Background())
	ns := &NS{st: st, exitContext: ctx, cancel: cancel}
	ns.dc = dc
	// start cache subscriber and listen for updates
	ns.groupCache = cache.NewScanGroupSubscriber(ctx, st)
	return ns
}

// Init the redisclient and dns client.
func (ns *NS) Init(config []byte) error {
	go ns.debug()
	// populate cache
	return nil
}

// Stop this module from running, and close down subscriptions
func (ns *NS) Stop(ctx context.Context) {
	ns.cancel()
}

// Name returns the module name
func (ns *NS) Name() string {
	return "NS"
}

func (ns *NS) debug() {
	stackTicker := time.NewTicker(time.Minute * 15)
	defer stackTicker.Stop()
	for {
		select {
		case <-ns.exitContext.Done():
			return
		case <-stackTicker.C:
			buf := make([]byte, 1<<20)
			stacklen := runtime.Stack(buf, true)
			log.Printf("*** goroutine dump...\n%s\n*** end\n", buf[:stacklen])
		}
	}
}

// Analyze an address, extracts NS, MX, A, AAAA, CNAME records
// TODO: add error if shutting down so dispatcher can retry
func (ns *NS) Analyze(ctx context.Context, userContext am.UserContext, address *am.ScanGroupAddress) (*am.ScanGroupAddress, map[string]*am.ScanGroupAddress, error) {
	logger := module.DefaultLogger(userContext, address)

	nsRecords := make(map[string]*am.ScanGroupAddress, 0)

	address.LastScannedTime = time.Now().UnixNano()
	if !ns.shouldAnalyze(address) {
		logger.Info().Msg("will not analyze")
		return address, nsRecords, nil
	}

	logger.Info().Msg("will analyze host")
	resolvedHosts := ns.analyzeHost(ctx, logger, address)
	logger.Info().Msg("will analyze ip")
	resolvedIPs := ns.analyzeIP(ctx, logger, address)

	module.AddAddressToMap(nsRecords, resolvedHosts)
	module.AddAddressToMap(nsRecords, resolvedIPs)

	if address.HostAddress == "" {
		return address, nsRecords, nil
	}

	etld, err := parsers.GetETLD(address.HostAddress)
	if err != nil || etld == "" {
		// push nsRecords
		return address, nsRecords, nil
	}

	ok, err := ns.st.DoNSRecords(ctx, address.OrgID, address.GroupID, nsExpire, etld)
	if err != nil {
		logger.Warn().Err(err).Str("etld", etld).Msg("unable to analyze ns records")
	}

	if ok {
		logger.Info().Msg("will analyze zone")
		zoneRecords := ns.analyzeZone(ctx, logger, etld, address)
		logger.Info().Int("zone_records", len(zoneRecords)).Str("etld", etld).Msg("got records")
		module.AddAddressToMap(nsRecords, zoneRecords)
	}

	// push nsRecords
	logger.Info().Msg("returning records")
	return address, nsRecords, nil
}

// shouldAnalyze determines if we should analyze the specific address or not
func (ns *NS) shouldAnalyze(address *am.ScanGroupAddress) bool {
	if address.IsHostedService || module.IsHostedDomain(address.HostAddress) {
		return false
	}

	switch uint16(address.NSRecord) {
	case dns.TypeMX, dns.TypeNS, dns.TypeSRV:
		return false
	}
	return true
}

// analyzeZone looks up various supporting records for a zone (mx/ns/axfr)
func (ns *NS) analyzeZone(ctx context.Context, logger zerolog.Logger, zone string, address *am.ScanGroupAddress) []*am.ScanGroupAddress {
	nsData := make([]*am.ScanGroupAddress, 0)

	logger.Info().Msg("will analyze mx")
	r, err := ns.dc.LookupMX(ctx, zone)
	if err == nil {
		for _, host := range r.Hosts {
			newAddress := module.NewAddressFromDNS(address, "", parsers.FQDNTrim(host), am.DiscoveryNSQueryOther, uint(r.RecordType))
			nsData = append(nsData, newAddress)
		}
	}
	logger.Info().Msg("will analyze ns")
	r, err = ns.dc.LookupNS(ctx, zone)
	if err == nil {
		for _, host := range r.Hosts {
			newAddress := module.NewAddressFromDNS(address, "", parsers.FQDNTrim(host), am.DiscoveryNSQueryOther, uint(r.RecordType))
			nsData = append(nsData, newAddress)
		}
	}
	logger.Info().Msg("will analyze axfr")
	axfr, err := ns.dc.DoAXFR(ctx, zone)
	if err != nil {
		return nsData
	}

	for _, result := range axfr {
		// TODO report axfr ns servers as a finding
		for _, r := range result {
			if len(r.Hosts) == 1 {
				for _, ip := range r.IPs {
					newAddress := module.NewAddressFromDNS(address, ip, parsers.FQDNTrim(r.Hosts[0]), am.DiscoveryNSAXFR, uint(r.RecordType))
					nsData = append(nsData, newAddress)
				}
			} else if len(r.IPs) == 1 {
				for _, host := range r.Hosts {
					newAddress := module.NewAddressFromDNS(address, r.IPs[0], parsers.FQDNTrim(host), am.DiscoveryNSAXFR, uint(r.RecordType))
					nsData = append(nsData, newAddress)
				}
			}
		}
	}
	return nsData
}

// analyzeIP for this address, finding potentially new hostnames
// if ip == same host, update last seen and scanned time
// if host == empty, add host to the first record
// if ip == different host, create new address
// if host address from address was not returned in any records, try to do a look up??
func (ns *NS) analyzeIP(ctx context.Context, logger zerolog.Logger, address *am.ScanGroupAddress) []*am.ScanGroupAddress {
	nsData := make([]*am.ScanGroupAddress, 0)

	if address.IPAddress == "" {
		return nsData
	}

	r, err := ns.dc.ResolveIP(ctx, address.IPAddress)
	if err != nil || r == nil {
		logger.Error().Err(err).Msg("unable to resolve ip")
		return nsData
	}

	foundOriginal := false

	// we may get multiple hosts back, so check if we've ever found it before?
	for _, host := range r.Hosts {
		// we've seen this same host before *or* never resolved this ip before
		if host == address.HostAddress || address.HostAddress == "" && !foundOriginal {
			foundOriginal = true
			address.HostAddress = parsers.FQDNTrim(host)
			if !address.IsHostedService && address.HostAddress != "" {
				address.IsHostedService = module.IsHostedDomain(address.HostAddress)
			}
			address.LastSeenTime = time.Now().UnixNano()
			address.NSRecord = int32(r.RecordType)
			// update the hash address now that we have a proper host for it
			address.AddressHash = convert.HashAddress(address.IPAddress, host)
			continue
		}
		// or we got a new hostname when attempting to resolve this ip.
		// Copy details from original address into the new address
		newAddress := module.NewAddressFromDNS(address, address.IPAddress, parsers.FQDNTrim(host), am.DiscoveryNSQueryIPToName, uint(r.RecordType))
		newAddress.ConfidenceScore = module.CalculateConfidence(logger, address, newAddress)
		nsData = append(nsData, newAddress)
	}

	return nsData
}

// analyzeHost resolves the host address to ips
func (ns *NS) analyzeHost(ctx context.Context, logger zerolog.Logger, address *am.ScanGroupAddress) []*am.ScanGroupAddress {

	nsData := make([]*am.ScanGroupAddress, 0)
	if address.HostAddress == "" {
		return nsData
	}

	r, err := ns.dc.ResolveName(ctx, address.HostAddress)
	if err != nil {
		logger.Error().Err(err).Msg("unable to resolve host")
		return nsData
	}

	// we don't need to test for original here because we just take the first one
	// (i == 0 && j == 0)
	for i, rr := range r {
		for j, ip := range rr.IPs {
			if i == 0 && j == 0 {
				if !address.IsHostedService && address.HostAddress != "" {
					address.IsHostedService = module.IsHostedDomain(address.HostAddress)
				}
				address.IPAddress = ip
				address.LastSeenTime = time.Now().UnixNano()
				address.NSRecord = int32(rr.RecordType)
				address.AddressHash = convert.HashAddress(ip, address.HostAddress)
				continue
			}

			newAddress := module.NewAddressFromDNS(address, ip, address.HostAddress, am.DiscoveryNSQueryNameToIP, uint(rr.RecordType))
			newAddress.ConfidenceScore = module.CalculateConfidence(logger, address, newAddress)
			nsData = append(nsData, newAddress)
		}
	}
	return nsData
}
