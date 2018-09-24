package mock

import (
	"context"

	"github.com/linkai-io/am/am"
	"github.com/linkai-io/am/pkg/redisclient"
)

type DispatcherState struct {
	InitFn        func(config []byte) error
	InitFnInvoked bool

	PopAddressesFn      func(ctx context.Context, userContext am.UserContext, scanGroupID int, limit int) (map[string]*am.ScanGroupAddress, error)
	PopAddressesInvoked bool

	SubscribeFn      func(ctx context.Context, onStartFn redisclient.SubOnStart, onMessageFn redisclient.SubOnMessage, channels ...string) error
	SubscribeInvoked bool

	GetGroupFn      func(ctx context.Context, orgID, scanGroupID int, wantModules bool) (*am.ScanGroup, error)
	GetGroupInvoked bool

	GroupStatusFn      func(ctx context.Context, userContext am.UserContext, scanGroupID int) (bool, am.GroupStatus, error)
	GroupStatusInvoked bool

	PutAddressesFn      func(ctx context.Context, userContext am.UserContext, scanGroupID int, addresses []*am.ScanGroupAddress) error
	PutAddressesInvoked bool

	PutAddressMapFn      func(ctx context.Context, userContext am.UserContext, scanGroupID int, addresses map[string]*am.ScanGroupAddress) error
	PutAddressMapInvoked bool

	FilterNewFn      func(ctx context.Context, orgID, scanGroupID int, addresses map[string]*am.ScanGroupAddress) (map[string]*am.ScanGroupAddress, error)
	FilterNewInvoked bool
}

func (s *DispatcherState) Init(config []byte) error {
	return nil
}

func (s *DispatcherState) PopAddresses(ctx context.Context, userContext am.UserContext, scanGroupID int, limit int) (map[string]*am.ScanGroupAddress, error) {
	s.PopAddressesInvoked = true
	return s.PopAddressesFn(ctx, userContext, scanGroupID, limit)
}

func (s *DispatcherState) Subscribe(ctx context.Context, onStartFn redisclient.SubOnStart, onMessageFn redisclient.SubOnMessage, channels ...string) error {
	s.SubscribeInvoked = true
	return s.SubscribeFn(ctx, onStartFn, onMessageFn, channels...)
}

func (s *DispatcherState) GetGroup(ctx context.Context, orgID int, scanGroupID int, wantModules bool) (*am.ScanGroup, error) {
	s.GetGroupInvoked = true
	return s.GetGroupFn(ctx, orgID, scanGroupID, wantModules)
}

func (s *DispatcherState) GroupStatus(ctx context.Context, userContext am.UserContext, scanGroupID int) (bool, am.GroupStatus, error) {
	s.GroupStatusInvoked = true
	return s.GroupStatusFn(ctx, userContext, scanGroupID)
}

func (s *DispatcherState) PutAddresses(ctx context.Context, userContext am.UserContext, scanGroupID int, addresses []*am.ScanGroupAddress) error {
	s.PutAddressesInvoked = true
	return s.PutAddressesFn(ctx, userContext, scanGroupID, addresses)
}

func (s *DispatcherState) PutAddressMap(ctx context.Context, userContext am.UserContext, scanGroupID int, addresses map[string]*am.ScanGroupAddress) error {
	s.PutAddressMapInvoked = true
	return s.PutAddressMapFn(ctx, userContext, scanGroupID, addresses)
}
func (s *DispatcherState) FilterNew(ctx context.Context, orgID, scanGroupID int, addresses map[string]*am.ScanGroupAddress) (map[string]*am.ScanGroupAddress, error) {
	s.FilterNewInvoked = true
	return s.FilterNewFn(ctx, orgID, scanGroupID, addresses)
}
