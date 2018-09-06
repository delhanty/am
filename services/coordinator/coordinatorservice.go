package coordinator

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/linkai-io/am/am"
	"github.com/linkai-io/am/services/coordinator/queue"
	"github.com/linkai-io/am/services/coordinator/state"
)

var (
	modules            = []string{"ns", "dnsbrute", "port", "web"}
	ErrScanGroupPaused = errors.New("scan group is currently paused")
)

// Service for interfacing with postgresql/rds
type Service struct {
	state           state.Stater
	queueClient     queue.Queue
	addressClient   am.AddressService
	scanGroupClient am.ScanGroupService
}

// New returns an empty Service
func New(state state.Stater, addressClient am.AddressService, scanGroupClient am.ScanGroupService, queueClient queue.Queue) *Service {
	return &Service{state: state, addressClient: addressClient, scanGroupClient: scanGroupClient, queueClient: queueClient}
}

// Init by parsing the config and initializing the database pool
func (s *Service) Init(config []byte) error {
	return nil
}

// StartGroup initializes state system and queues if they do not exist, or updates with scan group details
func (s *Service) StartGroup(ctx context.Context, userContext am.UserContext, scanGroupID int) error {
	var queueMap map[string]string

	oid, group, err := s.scanGroupClient.Get(ctx, userContext, scanGroupID)
	if err != nil {
		return err
	}

	if oid != userContext.GetOrgID() {
		return am.ErrOrgIDMismatch
	}

	if group.Paused {
		return ErrScanGroupPaused
	}

	exists, _, err := s.state.GroupStatus(ctx, userContext, scanGroupID)
	if err != nil {
		return err
	}

	if !exists {
		// create queues
		if queueMap, err = s.createGroupQueues(ctx, group); err != nil {
			return err
		}

		// update/create configuration
		if err := s.state.Put(ctx, userContext, group, queueMap); err != nil {
			return err
		}
	}

	wantModules := true
	cachedGroup, err := s.state.GetGroup(ctx, userContext.GetOrgID(), scanGroupID, wantModules)
	if err != nil {
		return err
	}
	// TODO: if config is paused but group is not, handle retrieveing s3 bucket dumped
	// messages
	if cachedGroup.ModifiedTime < group.ModifiedTime {
		if err := s.state.Delete(ctx, userContext, group.GroupID); err != nil {
			return err
		}

		// update/create configuration
		if err := s.state.Put(ctx, userContext, group, queueMap); err != nil {
			return err
		}
	}

	if err := s.pushAddresses(ctx, userContext, scanGroupID); err != nil {
		return err
	}

	return s.state.Start(ctx, userContext, group.GroupID)
}

// create queue for scan group for each module type and store queue urls in redis
// TODO: shard them? if > 120,000 addresses this won't work will need to create a group
// of queues for each.
func (s *Service) createGroupQueues(ctx context.Context, group *am.ScanGroup) (map[string]string, error) {
	key := fmt.Sprintf("%d_%d_", group.OrgID, group.GroupID)
	queueMap := make(map[string]string, 0)
	queueName := key + "input"
	queueURL, err := s.queueClient.Create(queueName)
	if err != nil {
		return nil, err
	}

	queueMap[queueName] = queueURL
	for _, module := range modules {
		queueName = key + module
		queueURL, err := s.queueClient.Create(queueName)
		if err != nil {
			return nil, err
		}
		queueMap[queueName] = queueURL
	}
	return queueMap, nil
}

func (s *Service) pushAddresses(ctx context.Context, userContext am.UserContext, scanGroupID int) error {
	now := time.Now()
	// TODO: do smart calculation on size of scan group addresses
	then := now.Add(time.Duration(-4) * time.Hour).UnixNano()
	filter := &am.ScanGroupAddressFilter{
		OrgID:               userContext.GetOrgID(),
		GroupID:             scanGroupID,
		Start:               0,
		Limit:               1000,
		WithLastScannedTime: true,
		SinceScannedTime:    then,
		WithIgnored:         true,
	}

	// push addresses to state
	for {
		_, addrs, err := s.addressClient.Get(ctx, userContext, filter)
		if err != nil {
			return err
		}

		if len(addrs) == 0 {
			break
		}
		// get last addressid and update start for filter.
		filter.Start = addrs[len(addrs)-1].AddressID
		if err := s.state.PushAddresses(ctx, userContext, addrs); err != nil {
			return err
		}

		if err := s.queueClient.PushAddresses(ctx, addrs); err != nil {
			return err
		}
	}

	return nil
}
