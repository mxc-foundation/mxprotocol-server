package gateway

import (
	"context"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api/clients/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
	"time"
)

func Setup() error {
	log.Info("Syncronize gateways from appserver")
	syncGatewaysFromAppserverByBatch()
	return nil
}

func syncGatewaysFromAppserverByBatch() {
	// get gateway list from local database

	// get gateway list from appserver

	// do synchronization

}

func SyncGatewayProfileByMacFromAppserver(gwId int64, mac string) error {
	client, err := appserver.GetPool().Get(config.Cstruct.AppServer.Server, []byte(config.Cstruct.AppServer.CACert),
		[]byte(config.Cstruct.AppServer.TLSCert), []byte(config.Cstruct.AppServer.TLSKey))
	if err != nil {
		return err
	}

	gateway, err := client.GetGatewayByMac(context.Background(), &api.GetGatewayByMacRequest{Mac: mac})
	if err == nil && gateway.GwProfile == nil {
		// gateway no longer exist, delete from database
		err := db.Gateway.SetGatewayMode(gwId, types.GW_DELETED)
		if err != nil {
			log.WithError(err).Warn("gateway/SyncGatewayProfileByMacFromAppserver: gwId", gwId)
		}
	} else if err == nil {
		// get gateway successfully, add/update device
		walletId, err := db.Wallet.GetWalletIdFromOrgId(gateway.OrgId)
		if err != nil {
			log.WithError(err).Error("gateway/SyncGatewayProfileByMacFromAppserver: gateway is not linked to any wallet")
			err := db.Gateway.SetGatewayMode(gwId, types.GW_DELETED)
			if err != nil {
				log.WithError(err).Warn("gateway/SyncGatewayProfileByMacFromAppserver: gwId", gwId)
			}
			// in this case, it is not necessary to retry again
			return nil
		}

		_, err = db.Gateway.InsertGateway(types.Gateway{
			Mac:         mac,
			FkWallet:    walletId,
			Mode:        types.GW_WHOLE_NETWORK,
			CreatedAt:   time.Now(),
			OrgId:       gateway.OrgId,
			Description: gateway.GwProfile.Description,
			Name:        gateway.GwProfile.Name,
		})
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
