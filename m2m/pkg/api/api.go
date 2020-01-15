package api

import (
	log "github.com/sirupsen/logrus"
	grpcAppserver "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_server"
	m2mServer "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"github.com/mxc-foundation/mxprotocol-server/m2m/api/networkserver"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/pkg/api/m2m_appserver"
	networkserverApi "github.com/mxc-foundation/mxprotocol-server/m2m/pkg/api/m2m_networkserver"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/config"
	"google.golang.org/grpc"
	"net"
	"time"
)

func SetupHTTPServer(conf config.MxpConfig) error {
	server := grpc.NewServer()

	// register all servers here
	grpcAppserver.RegisterM2MServerServiceServer(server, api.NewM2MServerAPI())

	m2mServer.RegisterDeviceServiceServer(server, api.NewM2MServerAPI())
	m2mServer.RegisterGatewayServiceServer(server, api.NewM2MServerAPI())
	m2mServer.RegisterMoneyServiceServer(server, api.NewM2MServerAPI())
	m2mServer.RegisterServerInfoServiceServer(server, api.NewM2MServerAPI())
	m2mServer.RegisterSettingsServiceServer(server, api.NewM2MServerAPI())
	m2mServer.RegisterStakingServiceServer(server, api.NewM2MServerAPI())
	m2mServer.RegisterSuperNodeServiceServer(server, api.NewM2MServerAPI())
	m2mServer.RegisterTopUpServiceServer(server, api.NewM2MServerAPI())
	m2mServer.RegisterWalletServiceServer(server, api.NewM2MServerAPI())
	m2mServer.RegisterWithdrawServiceServer(server, api.NewM2MServerAPI())

	networkserver.RegisterM2MServerServiceServer(server, networkserverApi.NewM2MNetworkServerAPI())

	ln, err := net.Listen("tcp", conf.M2MServer.HttpServer.Bind)
	if err != nil {
		log.WithError(err).Error("start api listener error")
		return err
	}
	go server.Serve(ln)

	// give the http server some time to start
	time.Sleep(time.Millisecond * 100)

	return nil
}
