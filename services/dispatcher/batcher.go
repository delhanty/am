package dispatcher

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/linkai-io/am/am"
	"github.com/rs/zerolog/log"
)

type Batcher struct {
	addressClient am.AddressService
	userContext   am.UserContext

	// for pooling
	batchCount int
	count      int32
	results    chan *am.ScanGroupAddress

	//
	doneCh chan struct{}
}

func NewBatcher(userContext am.UserContext, addressClient am.AddressService, batchCount int) *Batcher {
	return &Batcher{
		addressClient: addressClient,
		userContext:   userContext,
		batchCount:    batchCount,
	}
}

func (b *Batcher) Init() error {
	b.doneCh = make(chan struct{})
	b.results = make(chan *am.ScanGroupAddress, b.batchCount)
	go b.InsertBatch()
	return nil
}

func (b *Batcher) Add(result *am.ScanGroupAddress) {
	select {
	case b.results <- result:
		atomic.AddInt32(&b.count, 1)
	}
}

func (b *Batcher) Drain() map[string]*am.ScanGroupAddress {
	results := make(map[string]*am.ScanGroupAddress, 0)
	for {
		select {
		case result := <-b.results:
			results[result.AddressHash] = result
			atomic.AddInt32(&b.count, -1)
			if len(results) >= b.batchCount {
				log.Info().Int("count", len(results)).Msg("Uploader Drained")
				return results
			}
		default:
			return results
		}
	}
}

func (b *Batcher) Count() int32 {
	return atomic.LoadInt32(&b.count)
}

func (b *Batcher) InsertBatch() {
	t := time.NewTicker(time.Second * 1)
	defer t.Stop()
	for {
		select {
		case <-b.doneCh:
			addrs := b.Drain()
			b.update(addrs)
			return
		case <-t.C:
			addrs := b.Drain()
			b.update(addrs)
		}
	}
}

func (b *Batcher) Done() {
	addrs := b.Drain()
	b.update(addrs)
	close(b.doneCh)
}

func (b *Batcher) update(addresses map[string]*am.ScanGroupAddress) {
	var err error
	var count int

	if addresses == nil || len(addresses) == 0 {
		return
	}

	ctx := context.Background()

	_, count, err = b.addressClient.Update(ctx, b.userContext, addresses)
	if err != nil {
		log.Error().Err(err).Msg("Unable to insert batch of addresses")
		for k, v := range addresses {
			log.Info().Msgf("hash: %s val: %#v", k, v)
		}
		return
	}

	log.Info().Int("count", count).Msg("inserted addresses")
}
