package appserver

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"sync"
	"time"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"

)

var p Pool

type Pool interface {
	Get(hostname string, caCert, tlsCert, tlsKey []byte) (appserver.AppServerServiceClient, error)
}

type appserverServiceClient struct {
	client     appserver.AppServerServiceClient
	clientConn *grpc.ClientConn
	caCert     []byte
	tlsCert    []byte
	tlsKey     []byte
}

func Setup(conf config.MxpConfig) error {
	p = &pool{
		appserverServiceClients: make(map[string]appserverServiceClient),
	}
	return nil
}

func GetPool() Pool {
	return p
}

func SetPool(pp Pool) {
	p = pp
}

type pool struct {
	sync.RWMutex
	appserverServiceClients map[string]appserverServiceClient
}

// Get returns a AppServerServiceClient for the given server (hostname:ip).
func (p *pool) Get(hostname string, caCert, tlsCert, tlsKey []byte) (appserver.AppServerServiceClient, error) {
	defer p.Unlock()
	p.Lock()

	var connect bool
	c, ok := p.appserverServiceClients[hostname]
	if !ok {
		connect = true
	}

	// if the connection exists in the map, but when the certificates changed
	// try to cloe the connection and re-connect
	if ok && (!bytes.Equal(c.caCert, caCert) || !bytes.Equal(c.tlsCert, tlsCert) || !bytes.Equal(c.tlsKey, tlsKey)) {
		c.clientConn.Close()
		delete(p.appserverServiceClients, hostname)
		connect = true
	}

	if connect {
		clientConn, nsClient, err := p.createClient(hostname, caCert, tlsCert, tlsKey)
		if err != nil {
			return nil, errors.Wrap(err, "create m2m-server api client error")
		}
		c = appserverServiceClient{
			client:     nsClient,
			clientConn: clientConn,
			caCert:     caCert,
			tlsCert:    tlsCert,
			tlsKey:     tlsKey,
		}
		p.appserverServiceClients[hostname] = c
	}

	return c.client, nil
}

func (p *pool) createClient(hostname string, caCert, tlsCert, tlsKey []byte) (*grpc.ClientConn, appserver.AppServerServiceClient, error) {
	logrusEntry := log.NewEntry(log.StandardLogger())
	logrusOpts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel),
	}

	nsOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(
			grpc_logrus.UnaryClientInterceptor(logrusEntry, logrusOpts...),
		),
		grpc.WithStreamInterceptor(
			grpc_logrus.StreamClientInterceptor(logrusEntry, logrusOpts...),
		),
	}

	if len(caCert) == 0 && len(tlsCert) == 0 && len(tlsKey) == 0 {
		nsOpts = append(nsOpts, grpc.WithInsecure())
		log.WithField("server", hostname).Warning("creating insecure appserver client")
	} else {
		log.WithField("server", hostname).Info("creating appserver client")
		cert, err := tls.X509KeyPair(tlsCert, tlsKey)
		if err != nil {
			return nil, nil, errors.Wrap(err, "load x509 keypair error")
		}

		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			return nil, nil, errors.Wrap(err, "append ca cert to pool error")
		}

		nsOpts = append(nsOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		})))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	appserverServiceClient, err := grpc.DialContext(ctx, hostname, nsOpts...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "dial appserver service api error")
	}

	return appserverServiceClient, appserver.NewAppServerServiceClient(appserverServiceClient), nil
}
