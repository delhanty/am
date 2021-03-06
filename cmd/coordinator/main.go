package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/linkai-io/am/am"
	"github.com/linkai-io/am/pkg/initializers"
	"github.com/linkai-io/am/pkg/lb/consul"
	"github.com/linkai-io/am/pkg/metrics/load"
	"github.com/linkai-io/am/pkg/retrier"
	"github.com/linkai-io/am/pkg/secrets"
	"github.com/rs/zerolog"

	coordinatorprotoservice "github.com/linkai-io/am/protocservices/coordinator"
	"github.com/linkai-io/am/protocservices/metrics"
	"github.com/linkai-io/am/services/coordinator"
	coordprotoc "github.com/linkai-io/am/services/coordinator/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const (
	serviceKey = am.CoordinatorServiceKey
)

var (
	appConfig initializers.AppConfig
)

func init() {
	appConfig.Env = os.Getenv("APP_ENV")
	appConfig.Region = os.Getenv("APP_REGION")
	appConfig.SelfRegister = os.Getenv("APP_SELF_REGISTER")
	appConfig.Addr = os.Getenv("APP_ADDR")
	appConfig.ServiceKey = serviceKey
	consulAddr := initializers.ServiceDiscovery(&appConfig)
	consul.RegisterDefault(time.Second*5, consulAddr) // Address comes from CONSUL_HTTP_ADDR or from aws metadata
}

// main starts the CoordinatorService
func main() {
	var err error

	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("service", "CoordinatorService").Logger()

	sec := secrets.NewSecretsCache(appConfig.Env, appConfig.Region)

	systemOrgID, err := sec.SystemOrgID()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to get system org id")
	}

	systemUserID, err := sec.SystemUserID()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to get system user id")
	}

	if appConfig.Addr == "" {
		appConfig.Addr = ":50051"
	}

	listener, err := net.Listen("tcp", appConfig.Addr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	state := initializers.State(&appConfig)
	dispatcherClient := initializers.DispatcherClient()
	scanGroupClient := initializers.SGClient()
	orgClient := initializers.OrgClient()

	service := coordinator.New(state, dispatcherClient, orgClient, scanGroupClient, systemOrgID, systemUserID)
	err = retrier.Retry(func() error {
		return service.Init(nil)
	})

	if err != nil {
		log.Fatal().Err(err).Msg("initializing service failed")
	}

	s := grpc.NewServer()
	r := load.NewRateReporter(time.Minute)

	coordp := coordprotoc.New(service, r)
	coordinatorprotoservice.RegisterCoordinatorServer(s, coordp)
	healthgrpc.RegisterHealthServer(s, health.NewServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	metrics.RegisterLoadReportServer(s, r)

	// check if self register
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	initializers.Self(ctx, &appConfig)

	log.Info().Msg("Starting Service")
	if err := s.Serve(listener); err != nil {
		log.Fatal().Err(err).Msg("failed to serve grpc")
	}
}
