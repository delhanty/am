package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"

	"github.com/linkai-io/am/pkg/secrets"

	"github.com/linkai-io/am/pkg/autocertcache"
	"github.com/linkai-io/am/pkg/protocauth"

	"github.com/linkai-io/am/pkg/dnsclient"
	"github.com/linkai-io/am/pkg/metrics/load"
	"github.com/linkai-io/am/pkg/portscanner"

	"github.com/linkai-io/am/am"
	"github.com/linkai-io/am/protocservices/metrics"
	moduleservice "github.com/linkai-io/am/protocservices/module/portscan"
	"github.com/linkai-io/am/services/module/portscan"
	modulerprotoc "github.com/linkai-io/am/services/module/portscan/protoc"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/acme/autocert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var env = os.Getenv("APP_ENV")

var dnsServers = []string{"216.98.138.160:53", "216.98.128.160:53", "1.1.1.1:53", "1.0.0.1:53", "64.6.64.6:53", "77.88.8.8:53", "74.82.42.42:53", "8.8.4.4:53", "8.8.8.8:53"}

var (
	credentialPath string
	hostname       string
	certPath       string
	listenPort     string
)

const (
	serviceKey = am.PortScanModuleServiceKey
)

func init() {
	flag.StringVar(&hostname, "host", "scanner1.linkai.io", "hostname to use for serving files from")
	flag.StringVar(&certPath, "certs", "/opt/scanner/certs", "path to autocert cache")
	flag.StringVar(&credentialPath, "creds", "/opt/scanner/", "path to aws credentials")
	flag.StringVar(&listenPort, "port", ":50052", "port to listen on 50052 for prod, 50053 for dev")
}

// GetTLS reads the tls cache from the scanwebserver certs path
func GetTLS(host, cacheDir string) (*tls.Config, error) {
	if _, err := os.Stat(cacheDir + "/scanner1.linkai.io"); err != nil {
		return nil, err
	}

	manager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocertcache.GroupDirCache(cacheDir),
		HostPolicy: autocert.HostWhitelist(host),
	}
	return &tls.Config{GetCertificate: manager.GetCertificate}, nil
}

func getPortToken() []byte {
	type awsCredentials struct {
		ID  string `json:"id"`
		Key string `json:"key"`
	}
	creds := &awsCredentials{}

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.key", env, env))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open credentials file")
	}

	if err := json.Unmarshal(data, creds); err != nil {
		log.Fatal().Err(err).Msg("failed to unmarshal creds")
	}

	sec := secrets.NewSecretsCache(env, "us-east-1")
	sec.WithCredentials(creds.ID, creds.Key)

	portToken, err := sec.GetPortScanToken()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get port scan token from secrets")
	}

	if portToken == "" {
		log.Fatal().Err(err).Msg("token was empty")
	}

	return []byte(portToken)
}

func main() {
	flag.Parse()

	dnsClient := dnsclient.New(dnsServers, 3)

	tlsConfig, err := GetTLS(hostname, certPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get tls certificates")
	}

	listener, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	log.Info().Str("port", listenPort).Msg("listening on port")

	executor := portscanner.NewSocketClient(env)
	if err := executor.Init(nil); err != nil {
		log.Fatal().Err(err).Msg("failed to initialize portscanner socket client")
	}

	r := load.NewRateReporter(time.Minute)

	service := portscan.New(executor, dnsClient)
	if err := service.Init(nil); err != nil {
		log.Fatal().Err(err).Msg("failed to start port scan service")
	}

	servicep := modulerprotoc.New(service, r)

	creds := credentials.NewTLS(tlsConfig)
	authorizer := protocauth.New(getPortToken())

	s := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(protocauth.UnaryServerInterceptor(authorizer.Authenticate)), grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		MinTime:             time.Second * 30,
		PermitWithoutStream: true,
	}))

	moduleservice.RegisterPortScanModuleServer(s, servicep)
	healthgrpc.RegisterHealthServer(s, health.NewServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	metrics.RegisterLoadReportServer(s, r)

	log.Info().Msg("Starting Service")
	if err := s.Serve(listener); err != nil {
		log.Fatal().Err(err).Msg("failed to serve grpc")
	}
}
