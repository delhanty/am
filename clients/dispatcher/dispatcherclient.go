package dispatcher

import (
	"context"
	"time"

	"github.com/bsm/grpclb"
	balancerpb "github.com/bsm/grpclb/grpclb_balancer_v1"
	"github.com/linkai-io/am/am"
	"github.com/linkai-io/am/pkg/convert"
	"github.com/linkai-io/am/pkg/retrier"
	"github.com/rs/zerolog/log"

	service "github.com/linkai-io/am/protocservices/dispatcher"
	"google.golang.org/grpc"
)

type Client struct {
	client service.DispatcherClient
}

func New() *Client {
	return &Client{}
}

func (c *Client) Init(config []byte) error {
	balancer := grpc.RoundRobin(grpclb.NewResolver(&grpclb.Options{
		Address: string(config),
	}))

	conn, err := grpc.Dial(am.DispatcherServiceKey, grpc.WithInsecure(), grpc.WithBalancer(balancer))
	if err != nil {
		return err
	}
	go debug(string(config))
	c.client = service.NewDispatcherClient(conn)
	return nil
}

func debug(addr string) {
	for {
		time.Sleep(5 * time.Second)
		cc, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Error().Err(err).Msg("dispatcher client error dialing address")
			continue
		}
		defer cc.Close()

		bc := balancerpb.NewLoadBalancerClient(cc)
		resp, err := bc.Servers(context.Background(), &balancerpb.ServersRequest{
			Target: am.DispatcherServiceKey,
		})
		if err != nil {
			log.Error().Err(err).Msg("dispatcher client error in resp")
			continue
		}
		if len(resp.Servers) == 0 {
			log.Warn().Msg("No dispatcher servers found")
		}
	}
}

func (c *Client) PushAddresses(ctx context.Context, userContext am.UserContext, scanGroupID int) error {
	in := &service.PushRequest{
		UserContext: convert.DomainToUserContext(userContext),
		GroupID:     int32(scanGroupID),
	}

	return retrier.Retry(func() error {
		var err error
		_, err = c.client.PushAddresses(ctx, in)
		return err
	})

}
