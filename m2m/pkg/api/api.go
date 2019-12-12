package api

import (
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	grpcAppserver "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/networkserver"
	appserverApi "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api/m2m_appserver"
	networkserverApi "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api/m2m_networkserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

var (
	bind            string
	tlsCert         string
	tlsKey          string
	jwtSecret       string
	corsAllowOrigin string
)

func SetupHTTPServer(conf config.MxpConfig) error {
	bind = conf.M2MServer.HttpServer.Bind
	tlsCert = conf.M2MServer.HttpServer.TLSCert
	tlsKey = conf.M2MServer.HttpServer.TLSKey
	jwtSecret = conf.M2MServer.HttpServer.JWTSecret
	corsAllowOrigin = os.Getenv("APPSERVER")

	server := grpc.NewServer()

	// register all servers here
	grpcAppserver.RegisterM2MServerServiceServer(server, appserverApi.NewM2MServerAPI())
	networkserver.RegisterM2MServerServiceServer(server, networkserverApi.NewM2MNetworkServerAPI())

	var clientHttpHandler http.Handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if 2 == r.ProtoMajor && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			server.ServeHTTP(w, r)
		} else {
			if clientHttpHandler == nil {
				w.WriteHeader(http.StatusNotImplemented)
				return
			}

			if corsAllowOrigin != "" {
				w.Header().Set("Access-Control-Allow-Origin", corsAllowOrigin)
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Grpc-Metadata-Authorization")

				if r.Method == "OPTIONS" {
					return
				}
			}

			clientHttpHandler.ServeHTTP(w, r)
		}
	})

	go func() {
		log.WithFields(log.Fields{
			"bind":     bind,
			"tls-cert": tlsCert,
			"tls-key":  tlsKey,
		}).Info("pkg/api: starting m2m api")

		if tlsCert == "" || tlsKey == "" {
			log.Fatal(http.ListenAndServe(bind, h2c.NewHandler(handler, &http2.Server{})))
		} else {
			log.Fatal(http.ListenAndServeTLS(
				bind,
				tlsCert,
				tlsKey,
				h2c.NewHandler(handler, &http2.Server{}),
			))
		}
	}()

	// give the http server some time to start
	time.Sleep(time.Millisecond * 100)

	return nil
}
