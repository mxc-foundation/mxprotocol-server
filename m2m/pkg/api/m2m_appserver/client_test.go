package appserver

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"testing"
	"time"
)

type M2MServerAPITestSuite struct {
	suite.Suite
}

func TestM2MServerAPI(t *testing.T) {
	suite.Run(t, new(M2MServerAPITestSuite))
}

type m2mServiceClient struct {
	clientConn *grpc.ClientConn
	caCert     []byte
	tlsCert    []byte
	tlsKey     []byte
}

func Get(hostname string, caCert, tlsCert, tlsKey []byte) (*grpc.ClientConn, error) {

	clientConn, err := createClient(hostname, caCert, tlsCert, tlsKey)
	if err != nil {
		log.WithError(err).Error("create m2m-server api client error")
		return nil, err
	}
	c := m2mServiceClient{
		clientConn: clientConn,
		caCert:     caCert,
		tlsCert:    tlsCert,
		tlsKey:     tlsKey,
	}

	return c.clientConn, nil
}

func createClient(hostname string, caCert, tlsCert, tlsKey []byte) (*grpc.ClientConn, error) {
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
		log.WithField("server", hostname).Warning("creating insecure m2m-server device service client")
	} else {
		log.WithField("server", hostname).Info("creating m2m-server device service client")
		cert, err := tls.X509KeyPair(tlsCert, tlsKey)
		if err != nil {
			log.WithError(err).Error("load x509 keypair error")
			return nil, err
		}

		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			log.WithError(err).Error("append ca cert to pool error")
			return nil, err
		}

		nsOpts = append(nsOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		})))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	devServiceClient, err := grpc.DialContext(ctx, hostname, nsOpts...)
	if err != nil {
		log.WithError(err).Error("dial m2m-server device service api error", hostname)
		return nil, err
	}

	return devServiceClient, nil
}

func (ts *M2MServerAPITestSuite) TestM2MServerAPIs() {
	assert := require.New(ts.T())
	ctx := context.Background()

	clientConn, err := Get("localhost:4000", []byte(""), []byte(""), []byte(""))
	assert.Nil(err)

	ts.T().Run("Call WalletService APIs", func(t *testing.T) {
		assert := require.New(t)
		m2mClient := m2m_ui.NewWalletServiceClient(clientConn)
		getBalanceRes, err := m2mClient.GetWalletBalance(ctx, &m2m_ui.GetWalletBalanceRequest{OrgId: 0})
		assert.Nil(err)
		log.Info("GetWalletBalance ", getBalanceRes.Balance)

		/*		m2mClient.GetVmxcTxHistory()
				assert.Nil(err)

				m2mClient.GetWalletUsageHist()
				assert.Nil(err)

				m2mClient.GetDlPrice()
				assert.Nil(err)*/

	})

	/*ts.T().Run("Call DeviceService APIs", func(t *testing.T) {
		assert := require.New(t)
		m2mClient.GetDeviceList()
		assert.Nil(err)

		m2mClient.GetDeviceProfile()
		assert.Nil(err)

		m2mClient.GetDeviceHistory()
		assert.Nil(err)

		m2mClient.SetDeviceMode()
		assert.Nil(err)


	})

	ts.T().Run("Call MoneyService APIs", func(t *testing.T) {
		assert := require.New(t)
		m2mClient.ModifyMoneyAccount()
		assert.Nil(err)

		m2mClient.GetChangeMoneyAccountHistory()
		assert.Nil(err)

		m2mClient.GetActiveMoneyAccount()
		assert.Nil(err)

	})

	ts.T().Run("Call GatewayService APIs", func(t *testing.T) {
		assert := require.New(t)
		m2mClient.GetGatewayList
		assert.Nil(err)

		m2mClient.GetGatewayProfile
		assert.Nil(err)

		m2mClient.GetGatewayHistory
		assert.Nil(err)

		m2mClient.SetGatewayMode
		assert.Nil(err)

	})

	ts.T().Run("Call ServerInfoService APIs", func(t *testing.T) {
		assert := require.New(t)
		m2mClient.GetVersion
		assert.Nil(err)

	})

	ts.T().Run("Call SettingsService APIs", func(t *testing.T) {
		assert := require.New(t)
		m2mClient.GetSettings
		assert.Nil(err)

		m2mClient.ModifySettings
		assert.Nil(err)

	})

	ts.T().Run("Call StakingService APIs", func(t *testing.T) {
		assert := require.New(t)
		m2mClient.Stake
		assert.Nil(err)

		m2mClient.Unstake
		assert.Nil(err)

		m2mClient.GetActiveStakes
		assert.Nil(err)

		m2mClient.GetStakingHistory
		assert.Nil(err)

		m2mClient.GetStakingPercentage
		assert.Nil(err)


	})

	ts.T().Run("Call SuperNodeService APIs", func(t *testing.T) {
		assert := require.New(t)
		m2mClient.GetSuperNodeActiveMoneyAccount
		assert.Nil(err)

		m2mClient.AddSuperNodeMoneyAccount
		assert.Nil(err)

	})

	ts.T().Run("Call TopUpService APIs", func(t *testing.T) {
		assert := require.New(t)
		m2mClient.GetTransactionsHistory
		assert.Nil(err)

		m2mClient.GetIncome
		assert.Nil(err)

		m2mClient.GetTopUpHistory
		assert.Nil(err)

		m2mClient.GetTopUpDestination
		assert.Nil(err)

	})

	ts.T().Run("Call WithdrawService APIs", func(t *testing.T) {
		assert := require.New(t)
		m2mClient.GetWithdrawFee
		assert.Nil(err)

		m2mClient.WithdrawReq
		assert.Nil(err)

		m2mClient.GetWithdrawHistory
		assert.Nil(err)

		m2mClient.ModifyWithdrawFee
		assert.Nil(err)

	})*/

}
