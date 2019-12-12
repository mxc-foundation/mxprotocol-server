package api

import (
	log "github.com/sirupsen/logrus"
	grpcAppserver "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/networkserver"
	appserverApi "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api/m2m_appserver"
	networkserverApi "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api/m2m_networkserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"google.golang.org/grpc"
	"net"
	"time"
)

func SetupHTTPServer(conf config.MxpConfig) error {
	server := grpc.NewServer()

	// register all servers here
	grpcAppserver.RegisterM2MServerServiceServer(server, appserverApi.NewM2MServerAPI())
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
