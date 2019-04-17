package mock

import (
	"context"

	"github.com/linkai-io/am/am"
)

type ModuleService struct {
	InitFn        func(config []byte) error
	InitFnInvoked bool

	AnalyzeFn      func(ctx context.Context, userContext am.UserContext, address *am.ScanGroupAddress) (*am.ScanGroupAddress, map[string]*am.ScanGroupAddress, error)
	AnalyzeInvoked bool
}

func (s *ModuleService) Init(config []byte) error {
	return nil
}

func (s *ModuleService) Analyze(ctx context.Context, userContext am.UserContext, address *am.ScanGroupAddress) (*am.ScanGroupAddress, map[string]*am.ScanGroupAddress, error) {
	s.AnalyzeInvoked = true
	return s.AnalyzeFn(ctx, userContext, address)
}
