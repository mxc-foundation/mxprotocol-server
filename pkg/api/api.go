package api

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"mxprotocol-server/api"
	"mxprotocol-server/pkg/config"
	"mxprotocol-server/pkg/services/withdraw"
	"net"
)

var bind string

func Setup(conf config.MxpConfig) error {
	bind = ":" + conf.General.BindPort

	log.WithFields(log.Fields{
		"bind": bind,
	}).Info("pkg/api: starting mxp-server api")

	server := grpc.NewServer()
	// register all servers here
	api.RegisterWithdrawServiceServer(server, withdraw.NewWithdrawServerAPI())
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		return errors.Wrap(err, "start mxp-server api listener error")
	}
	go server.Serve(lis)

	return nil
}
