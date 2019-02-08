package bigdata

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/linkai-io/am/pkg/initializers"

	"github.com/linkai-io/am/amtest"

	"github.com/linkai-io/am/am"
	"github.com/linkai-io/am/mock"
	"github.com/linkai-io/am/pkg/convert"
	"github.com/linkai-io/am/pkg/dnsclient"
)

func TestBigDataSubdomainsFirstRun(t *testing.T) {
	dc := dnsclient.New([]string{"1.1.1.1:53"}, 2)
	st := amtest.MockBigDataState()
	bds := &mock.BigDataService{}
	bds.GetCTSubdomainsFn = func(ctx context.Context, userContext am.UserContext, etld string) (time.Time, map[string]*am.CTSubdomain, error) {
		return time.Now(), nil, nil
	}

	bds.AddCTSubdomainsFn = func(ctx context.Context, userContext am.UserContext, etld string, queryTime time.Time, subdomains map[string]*am.CTSubdomain) error {
		t.Logf("adding records")
		return nil
	}

	bq := &mock.BigQuerier{}
	bq.QuerySubdomainsFn = func(ctx context.Context, from time.Time, etld string) (map[string]*am.CTSubdomain, error) {
		res := make(map[string]*am.CTSubdomain, 1)
		res["dev.console.linkai.io"] = &am.CTSubdomain{ETLD: etld, Subdomain: "dev.console.linkai.io"}
		return res, nil
	}

	bd := New(dc, st, bds, bq)
	ctx := context.Background()
	userContext := amtest.CreateUserContext(1, 1)
	address := testBuildAddress("1.1.1.1", "blah.linkai.io")

	_, newAddrs, err := bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}

	if bq.QuerySubdomainsInvoked == false {
		t.Fatal("query etld should have been invoked")
	}

	if len(newAddrs) == 0 {
		t.Fatalf("failed to find at least 1 new addresses in big data, got %d\n", len(newAddrs))
	}

}

func TestBigDataSubdomainsFirstRunInfra(t *testing.T) {
	if os.Getenv("INFRA_TESTS") == "" {
		t.Skip("skipping infrastructure tests")
	}

	dc := dnsclient.New([]string{"1.1.1.1:53"}, 2)
	st := initializers.State(&initializers.AppConfig{Env: "local", Region: "us-east-1"})
	bds := &mock.BigDataService{}
	bds.GetCTSubdomainsFn = func(ctx context.Context, userContext am.UserContext, etld string) (time.Time, map[string]*am.CTSubdomain, error) {
		return time.Now(), nil, nil
	}

	bds.AddCTSubdomainsFn = func(ctx context.Context, userContext am.UserContext, etld string, queryTime time.Time, subdomains map[string]*am.CTSubdomain) error {
		t.Logf("adding records")
		return nil
	}

	bq := &mock.BigQuerier{}
	bq.QuerySubdomainsFn = func(ctx context.Context, from time.Time, etld string) (map[string]*am.CTSubdomain, error) {
		res := make(map[string]*am.CTSubdomain, 1)
		res["dev.console.linkai.io"] = &am.CTSubdomain{ETLD: etld, Subdomain: "dev.console.linkai.io"}
		return res, nil
	}

	bd := New(dc, st, bds, bq)
	ctx := context.Background()
	userContext := amtest.CreateUserContext(1, 1)
	address := testBuildAddress("1.1.1.1", "blah.linkai.io")

	_, newAddrs, err := bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}

	if bq.QuerySubdomainsInvoked == false {
		t.Fatal("query etld should have been invoked")
	}

	if len(newAddrs) == 0 {
		t.Fatalf("failed to find at least 1 new addresses in big data, got %d\n", len(newAddrs))
	}
}

func TestBigDataSubdomainsAppendNewResults(t *testing.T) {
	dc := dnsclient.New([]string{"1.1.1.1:53"}, 2)
	st := amtest.MockBigDataState()
	bds := &mock.BigDataService{}
	bds.GetCTSubdomainsFn = func(ctx context.Context, userContext am.UserContext, etld string) (time.Time, map[string]*am.CTSubdomain, error) {
		res := make(map[string]*am.CTSubdomain, 1)
		res["dev.console.linkai.io"] = &am.CTSubdomain{ETLD: etld, Subdomain: "dev.console.linkai.io"}
		return time.Now().Add(time.Hour * -5), res, nil
	}

	bds.AddCTSubdomainsFn = func(ctx context.Context, userContext am.UserContext, etld string, queryTime time.Time, subdomains map[string]*am.CTSubdomain) error {
		t.Logf("adding records")
		return nil
	}

	bq := &mock.BigQuerier{}
	bq.QuerySubdomainsFn = func(ctx context.Context, from time.Time, etld string) (map[string]*am.CTSubdomain, error) {
		res := make(map[string]*am.CTSubdomain, 1)
		res["prod.console.linkai.io"] = &am.CTSubdomain{ETLD: etld, Subdomain: "prod.console.linkai.io"}
		return res, nil
	}

	bd := New(dc, st, bds, bq)
	ctx := context.Background()
	userContext := amtest.CreateUserContext(1, 1)
	address := testBuildAddress("1.1.1.1", "blah.linkai.io")

	_, newAddrs, err := bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}

	if bq.QuerySubdomainsInvoked == false {
		t.Fatal("query etld should have been invoked")
	}

	if len(newAddrs) == 0 {
		t.Fatalf("failed to find at least 1 new addresses in big data, got %d\n", len(newAddrs))
	}

}

func TestBigDataSubdomainsRerun(t *testing.T) {
	dc := dnsclient.New([]string{"1.1.1.1:53"}, 2)
	st := amtest.MockBigDataState()
	bds := &mock.BigDataService{}
	bds.GetCTSubdomainsFn = func(ctx context.Context, userContext am.UserContext, etld string) (time.Time, map[string]*am.CTSubdomain, error) {
		return time.Now(), nil, nil
	}

	bds.AddCTSubdomainsFn = func(ctx context.Context, userContext am.UserContext, etld string, queryTime time.Time, subdomains map[string]*am.CTSubdomain) error {
		t.Logf("adding records")
		return nil
	}

	bq := &mock.BigQuerier{}
	bq.QuerySubdomainsFn = func(ctx context.Context, from time.Time, etld string) (map[string]*am.CTSubdomain, error) {
		res := make(map[string]*am.CTSubdomain, 1)
		res["dev.console.linkai.io"] = &am.CTSubdomain{ETLD: etld, Subdomain: "dev.console.linkai.io"}
		return res, nil
	}

	bd := New(dc, st, bds, bq)
	ctx := context.Background()
	userContext := amtest.CreateUserContext(1, 1)
	address := testBuildAddress("1.1.1.1", "blah.linkai.io")

	_, newAddrs, err := bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}

	if bq.QuerySubdomainsInvoked == false {
		t.Fatal("query etld should have been invoked")
	}

	if len(newAddrs) == 0 {
		t.Fatalf("failed to find at least 1 new addresses in big data, got %d\n", len(newAddrs))
	}

	// reset and re-run analysis
	bq.QuerySubdomainsInvoked = false
	bds.GetCTSubdomainsFn = func(ctx context.Context, userContext am.UserContext, etld string) (time.Time, map[string]*am.CTSubdomain, error) {
		res := make(map[string]*am.CTSubdomain, 1)
		res["dev.console.linkai.io"] = &am.CTSubdomain{ETLD: etld, Subdomain: "dev.console.linkai.io"}
		return time.Now(), res, nil
	}

	_, newAddrs, err = bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}

	if bq.QuerySubdomainsInvoked == true {
		t.Fatalf("re-run should not have invoked query subdomains")
	}
}

func TestBigDataFirstRun(t *testing.T) {
	// re-enable once we have money
	t.Skip()
	dc := dnsclient.New([]string{"1.1.1.1:53"}, 2)
	st := amtest.MockBigDataState()
	bds := &mock.BigDataService{}
	bds.GetCTFn = func(ctx context.Context, userContext am.UserContext, etld string) (time.Time, map[string]*am.CTRecord, error) {
		return time.Now(), nil, nil
	}
	bds.AddCTFn = func(ctx context.Context, userContext am.UserContext, etld string, queryTime time.Time, ctRecords map[string]*am.CTRecord) error {
		t.Logf("Adding records")
		return nil
	}

	bq := &mock.BigQuerier{}
	bq.QueryETLDFn = func(ctx context.Context, from time.Time, etld string) (map[string]*am.CTRecord, error) {
		return amtest.BuildCTRecords(etld, time.Now().UnixNano(), 1), nil
	}
	bd := New(dc, st, bds, bq)
	ctx := context.Background()
	userContext := amtest.CreateUserContext(1, 1)
	address := testBuildAddress("1.1.1.1", "blah.example.com")

	_, newAddrs, err := bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}

	if bq.QueryETLDInvoked == false {
		t.Fatal("query etld should have been invoked")
	}

	if len(newAddrs) != 2 {
		t.Fatalf("failed to find 2 new addresses in big data, got %d\n", len(newAddrs))
	}

}

func TestBigDataRerun(t *testing.T) {
	// re-enable once we have money
	t.Skip()
	dc := dnsclient.New([]string{"1.1.1.1:53"}, 2)
	st := amtest.MockBigDataState()
	bds := &mock.BigDataService{}

	// sets lastQueryTime to 1 day ago
	bds.GetCTFn = func(ctx context.Context, userContext am.UserContext, etld string) (time.Time, map[string]*am.CTRecord, error) {
		return time.Now().AddDate(0, 0, -1), amtest.BuildCTRecords(etld, time.Now().UnixNano(), 1), nil
	}

	bds.AddCTFn = func(ctx context.Context, userContext am.UserContext, etld string, queryTime time.Time, ctRecords map[string]*am.CTRecord) error {
		t.Logf("Adding records")
		return nil
	}

	bq := &mock.BigQuerier{}
	bq.QueryETLDFn = func(ctx context.Context, from time.Time, etld string) (map[string]*am.CTRecord, error) {
		return amtest.BuildCTRecords(etld, time.Now().UnixNano(), 2), nil
	}
	bd := New(dc, st, bds, bq)
	ctx := context.Background()
	userContext := amtest.CreateUserContext(1, 1)
	address := testBuildAddress("1.1.1.1", "blah.example.com")

	_, _, err := bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}

	if bq.QueryETLDInvoked == false {
		t.Fatal("query etld should have been invoked")
	}
}

func TestBigDataNoNewRecords(t *testing.T) {
	// re-enable once we have money
	t.Skip()
	dc := dnsclient.New([]string{"1.1.1.1:53"}, 2)
	st := amtest.MockBigDataState()
	bds := &mock.BigDataService{}

	// sets lastQueryTime to 1 day ago
	bds.GetCTFn = func(ctx context.Context, userContext am.UserContext, etld string) (time.Time, map[string]*am.CTRecord, error) {
		return time.Now().AddDate(0, 0, -1), amtest.BuildCTRecords(etld, time.Now().UnixNano(), 1), nil
	}

	bds.AddCTFn = func(ctx context.Context, userContext am.UserContext, etld string, queryTime time.Time, ctRecords map[string]*am.CTRecord) error {
		t.Logf("Adding records")
		return nil
	}

	bq := &mock.BigQuerier{}
	bq.QueryETLDFn = func(ctx context.Context, from time.Time, etld string) (map[string]*am.CTRecord, error) {
		return amtest.BuildCTRecords(etld, time.Now().UnixNano(), 1), nil
	}
	bd := New(dc, st, bds, bq)
	ctx := context.Background()
	userContext := amtest.CreateUserContext(1, 1)
	address := testBuildAddress("1.1.1.1", "blah.example.com")

	_, _, err := bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}

	if bq.QueryETLDInvoked == false {
		t.Fatal("query etld should have been invoked")
	}

	if bds.AddCTInvoked == true {
		t.Fatal("AddCT should not have been invoked since there are no new records.")
	}
}

func TestBigDataCacheTime(t *testing.T) {
	// re-enable once we have money
	t.Skip()
	dc := dnsclient.New([]string{"1.1.1.1:53"}, 2)
	st := amtest.MockBigDataState()
	bds := &mock.BigDataService{}

	// sets lastQueryTime to 1 day ago
	bds.GetCTFn = func(ctx context.Context, userContext am.UserContext, etld string) (time.Time, map[string]*am.CTRecord, error) {
		return time.Now(), amtest.BuildCTRecords(etld, time.Now().UnixNano(), 1), nil
	}

	bds.AddCTFn = func(ctx context.Context, userContext am.UserContext, etld string, queryTime time.Time, ctRecords map[string]*am.CTRecord) error {
		t.Logf("Adding records")
		return nil
	}

	bq := &mock.BigQuerier{}
	bq.QueryETLDFn = func(ctx context.Context, from time.Time, etld string) (map[string]*am.CTRecord, error) {
		return amtest.BuildCTRecords(etld, time.Now().UnixNano(), 1), nil
	}
	bd := New(dc, st, bds, bq)
	ctx := context.Background()
	userContext := amtest.CreateUserContext(1, 1)
	address := testBuildAddress("1.1.1.1", "blah.example.com")

	_, _, err := bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}

	if bq.QueryETLDInvoked == true {
		t.Fatal("query etld should not have been invoked")
	}

	if bds.AddCTInvoked == true {
		t.Fatal("AddCT should not have been invoked since there are no new records.")
	}
}

func TestBigDataDoCTTime(t *testing.T) {
	// re-enable once we have money
	t.Skip()
	dc := dnsclient.New([]string{"1.1.1.1:53"}, 2)
	st := amtest.MockBigDataState()
	bds := &mock.BigDataService{}

	// sets lastQueryTime to 1 day ago
	bds.GetCTFn = func(ctx context.Context, userContext am.UserContext, etld string) (time.Time, map[string]*am.CTRecord, error) {
		return time.Now(), amtest.BuildCTRecords(etld, time.Now().UnixNano(), 1), nil
	}

	bds.AddCTFn = func(ctx context.Context, userContext am.UserContext, etld string, queryTime time.Time, ctRecords map[string]*am.CTRecord) error {
		t.Logf("Adding records")
		return nil
	}

	bq := &mock.BigQuerier{}
	bq.QueryETLDFn = func(ctx context.Context, from time.Time, etld string) (map[string]*am.CTRecord, error) {
		return amtest.BuildCTRecords(etld, time.Now().UnixNano(), 1), nil
	}
	bd := New(dc, st, bds, bq)
	ctx := context.Background()
	userContext := amtest.CreateUserContext(1, 1)
	address := testBuildAddress("1.1.1.1", "blah.example.com")

	_, _, err := bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}
	_, _, err = bd.Analyze(ctx, userContext, address)
	if err != nil {
		t.Fatalf("failed to analyze using big data: %#v\n", err)
	}
}

func testBuildAddress(ip, host string) *am.ScanGroupAddress {
	addrHash := convert.HashAddress(ip, host)
	return &am.ScanGroupAddress{
		AddressID:           1,
		OrgID:               1,
		GroupID:             1,
		HostAddress:         host,
		IPAddress:           ip,
		DiscoveryTime:       0,
		DiscoveredBy:        "",
		LastScannedTime:     0,
		LastSeenTime:        0,
		ConfidenceScore:     100.0,
		UserConfidenceScore: 0.0,
		IsSOA:               false,
		IsWildcardZone:      false,
		IsHostedService:     false,
		Ignored:             false,
		FoundFrom:           "input_list",
		NSRecord:            0,
		AddressHash:         addrHash,
	}
}
