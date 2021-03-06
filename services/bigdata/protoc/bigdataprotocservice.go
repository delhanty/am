package protoc

import (
	"context"
	"time"

	"github.com/linkai-io/am/pkg/metrics/load"

	"github.com/linkai-io/am/pkg/convert"
	"github.com/linkai-io/am/protocservices/bigdata"

	"github.com/linkai-io/am/am"
)

type BigDataProtocService struct {
	bds      am.BigDataService
	reporter *load.RateReporter
}

func New(implementation am.BigDataService, reporter *load.RateReporter) *BigDataProtocService {
	return &BigDataProtocService{bds: implementation, reporter: reporter}
}

func (s *BigDataProtocService) GetCT(ctx context.Context, in *bigdata.GetCTRequest) (*bigdata.GetCTResponse, error) {
	s.reporter.Increment(1)
	ts, records, err := s.bds.GetCT(ctx, convert.UserContextToDomain(in.UserContext), in.ETLD)
	s.reporter.Increment(-1)
	if err != nil {
		return nil, err
	}

	return &bigdata.GetCTResponse{Time: ts.UnixNano(), Records: convert.DomainToCTRecords(records)}, nil
}

func (s *BigDataProtocService) AddCT(ctx context.Context, in *bigdata.AddCTRequest) (*bigdata.CTAddedResponse, error) {
	s.reporter.Increment(1)
	err := s.bds.AddCT(ctx, convert.UserContextToDomain(in.UserContext), in.ETLD, time.Unix(0, in.QueryTime), convert.CTRecordsToDomain(in.Records))
	s.reporter.Increment(-1)
	if err != nil {
		return nil, err
	}

	return &bigdata.CTAddedResponse{}, nil
}

func (s *BigDataProtocService) DeleteCT(ctx context.Context, in *bigdata.DeleteCTRequest) (*bigdata.CTDeletedResponse, error) {
	s.reporter.Increment(1)
	err := s.bds.DeleteCT(ctx, convert.UserContextToDomain(in.UserContext), in.ETLD)
	s.reporter.Increment(-1)
	if err != nil {
		return nil, err
	}

	return &bigdata.CTDeletedResponse{}, nil
}

func (s *BigDataProtocService) GetETLDs(ctx context.Context, in *bigdata.GetETLDsRequest) (*bigdata.GetETLDsResponse, error) {
	s.reporter.Increment(1)
	records, err := s.bds.GetETLDs(ctx, convert.UserContextToDomain(in.UserContext))
	s.reporter.Increment(-1)
	if err != nil {
		return nil, err
	}

	return &bigdata.GetETLDsResponse{ETLDs: convert.DomainToCTETLDs(records)}, nil
}

func (s *BigDataProtocService) GetCTSubdomains(ctx context.Context, in *bigdata.GetCTSubdomainsRequest) (*bigdata.GetCTSubdomainsResponse, error) {
	s.reporter.Increment(1)
	ts, records, err := s.bds.GetCTSubdomains(ctx, convert.UserContextToDomain(in.UserContext), in.ETLD)
	s.reporter.Increment(-1)
	if err != nil {
		return nil, err
	}

	return &bigdata.GetCTSubdomainsResponse{Time: ts.UnixNano(), Records: convert.DomainToCTSubdomainRecords(records)}, nil
}

func (s *BigDataProtocService) AddCTSubdomains(ctx context.Context, in *bigdata.AddCTSubdomainsRequest) (*bigdata.CTSubdomainsAddedResponse, error) {
	s.reporter.Increment(1)
	err := s.bds.AddCTSubdomains(ctx, convert.UserContextToDomain(in.UserContext), in.ETLD, time.Unix(0, in.QueryTime), convert.CTSubdomainRecordsToDomain(in.Records))
	s.reporter.Increment(-1)
	if err != nil {
		return nil, err
	}

	return &bigdata.CTSubdomainsAddedResponse{}, nil
}

func (s *BigDataProtocService) DeleteCTSubdomains(ctx context.Context, in *bigdata.DeleteCTSubdomainsRequest) (*bigdata.CTSubdomainsDeletedResponse, error) {
	s.reporter.Increment(1)
	err := s.bds.DeleteCTSubdomains(ctx, convert.UserContextToDomain(in.UserContext), in.ETLD)
	s.reporter.Increment(-1)
	if err != nil {
		return nil, err
	}

	return &bigdata.CTSubdomainsDeletedResponse{}, nil
}
