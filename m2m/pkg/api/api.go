package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	appserver_grpc "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	m2m_grpc "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m_ui"
	appserver_api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api/m2m_appserver"
	m2m_api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api/m2m_ui"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/static"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	bind            string
	tlsCert         string
	tlsKey          string
	jwtSecret       string
	corsAllowOrigin string
)

func SetupHTTPServer(conf config.MxpConfig) error {
	bind = conf.ApplicationServer.HttpServer.Bind
	tlsCert = conf.ApplicationServer.HttpServer.TLSCert
	tlsKey = conf.ApplicationServer.HttpServer.TLSKey
	jwtSecret = conf.ApplicationServer.HttpServer.JWTSecret
	corsAllowOrigin = conf.ApplicationServer.HttpServer.CORSAllowOrigin

	server := grpc.NewServer()

	// register all servers here
	m2m_grpc.RegisterWithdrawServiceServer(server, m2m_api.NewWithdrawServerAPI())
	m2m_grpc.RegisterMoneyServiceServer(server, m2m_api.NewMoneyServerAPI())
	m2m_grpc.RegisterTopUpServiceServer(server, m2m_api.NewTopUpServerAPI())
	m2m_grpc.RegisterWalletServiceServer(server, m2m_api.NewWalletServerAPI())
	m2m_grpc.RegisterSuperNodeServiceServer(server, m2m_api.NewSupernodeServerAPI())
	m2m_grpc.RegisterInternalServiceServer(server, auth.NewInternalServerAPI())
	m2m_grpc.RegisterDeviceServiceServer(server, m2m_api.NewDeviceServerAPI())
	m2m_grpc.RegisterGatewayServiceServer(server, m2m_api.NewGatewayServerAPI())
	appserver_grpc.RegisterM2MServerServiceServer(server, appserver_api.NewM2MServerAPI())

	var clientHttpHandler http.Handler
	var err error
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

	// setup the HTTP handler
	clientHttpHandler, err = setupHTTPAPI(conf)
	if err != nil {
		return err
	}

	return nil
}

func setupHTTPAPI(conf config.MxpConfig) (http.Handler, error) {
	r := mux.NewRouter()

	// setup json api handler
	jsonHandler, err := getJSONGateway(context.Background())
	if err != nil {
		return nil, err
	}

	log.WithField("path", "/api").Info("api/external: registering rest api handler and documentation endpoint")
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		data, err := static.Asset("swagger/index.html")
		if err != nil {
			log.WithError(err).Error("get swagger template error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}).Methods("get")
	r.PathPrefix("/api").Handler(jsonHandler)

	// setup static file server
	r.PathPrefix("/").Handler(http.FileServer(&assetfs.AssetFS{
		Asset:    static.Asset,
		AssetDir: static.AssetDir,
		Prefix:   "",
	}))

	return wsproxy.WebsocketProxy(r), nil
}

func getJSONGateway(ctx context.Context) (http.Handler, error) {
	// dial options for the grpc-gateway
	var grpcDialOpts []grpc.DialOption

	if tlsCert == "" || tlsKey == "" {
		grpcDialOpts = append(grpcDialOpts, grpc.WithInsecure())
	} else {
		b, err := ioutil.ReadFile(tlsCert)
		if err != nil {
			return nil, errors.Wrap(err, "read external api tls cert error")
		}
		cp := x509.NewCertPool()
		if !cp.AppendCertsFromPEM(b) {
			return nil, errors.Wrap(err, "failed to append certificate")
		}
		grpcDialOpts = append(grpcDialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			// given the grpc-gateway is always connecting to localhost, does
			// InsecureSkipVerify=true cause any security issues?
			InsecureSkipVerify: true,
			RootCAs:            cp,
		})))
	}

	bindParts := strings.SplitN(bind, ":", 2)
	if len(bindParts) != 2 {
		log.Fatal("get port from bind failed")
	}
	apiEndpoint := fmt.Sprintf("localhost:%s", bindParts[1])

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			EnumsAsInts:  false,
			EmitDefaults: true,
		},
	))

	// register all services
	if err := m2m_grpc.RegisterWithdrawServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register withdraw handler error")
	}
	if err := m2m_grpc.RegisterMoneyServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register ext_account handler error")
	}
	if err := m2m_grpc.RegisterTopUpServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register top_up handler error")
	}
	if err := m2m_grpc.RegisterWalletServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register top_up handler error")
	}
	if err := m2m_grpc.RegisterSuperNodeServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register top_up handler error")
	}
	if err := m2m_grpc.RegisterInternalServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register auth get jwt handler error")
	}

	return mux, nil
}
