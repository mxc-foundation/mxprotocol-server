package gateway

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api/clients/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
	"time"
)

var timer *time.Timer

func Setup() error {
	log.Info("Syncronize gateways from appserver")
	timer = time.AfterFunc(1*time.Second, syncGatewaysFromAppserverByBatch)

	// give it time to sync before whole service starts
	time.Sleep(5 * time.Second)
	return nil
}

func syncGatewaysFromAppserverByBatch() {
	// get gateway list from local database
	localGatewayList, err := db.Gateway.GetAllGateways()
	if err != nil {
		log.WithError(err).Error("service/device/syncGatewaysFromAppserverByBatch")
		// reset timer
		timer.Reset(10 * time.Second)
		return
	}
	log.Debug("syncGatewaysFromAppserverByBatch_local: count=", len(localGatewayList))

	// get gateway list from appserver
	client, err := appserver.GetPool().Get(config.Cstruct.AppServer.Server, []byte(config.Cstruct.AppServer.CACert),
		[]byte(config.Cstruct.AppServer.TLSCert), []byte(config.Cstruct.AppServer.TLSKey))
	if err != nil {
		log.WithError(err).Error("service/device/syncGatewaysFromAppserverByBatch")
		// reset timer
		timer.Reset(10 * time.Second)
		return
	}

	gwMacList, err := client.GetGatewayMacList(context.Background(), &empty.Empty{})
	if err != nil {
		log.WithError(err).Error("service/device/syncGatewaysFromAppserverByBatch")
		// reset timer
		timer.Reset(10 * time.Second)
		return
	}

	log.Debug("syncGatewaysFromAppserverByBatch_appserver: count=", len(gwMacList.GatewayMac), " list=", gwMacList.GatewayMac)

	// if len(localGatewayList) == 0, len(gwMacList.GatewayMac) == 0, just return
	if len(localGatewayList) == 0 && len(gwMacList.GatewayMac) == 0 {
		return
	}

	// if len(localGatewayList) == 0, len(gwMacList.GatewayMac) != 0, just insert new gateway
	if len(localGatewayList) == 0 && len(gwMacList.GatewayMac) != 0 {
		for _, v := range gwMacList.GatewayMac {
			gateway, err := getGatewayFromAppserver(v)
			if err != nil {
				log.WithError(err).Error("service/device/syncGatewaysFromAppserverByBatch")
				// reset timer
				timer.Reset(10 * time.Second)
				return
			}

			_, err = db.Gateway.InsertGateway(gateway)
			if err != nil {
				log.WithError(err).Error("service/device/syncGatewaysFromAppserverByBatch")
				timer.Reset(10 * time.Second)
				return
			}
		}

		return
	}

	//

	// do synchronization

}

func getGatewayFromAppserver(mac string) (types.Gateway, error) {
	gateway := types.Gateway{}
	appserverClient, err := appserver.GetPool().Get(config.Cstruct.AppServer.Server, []byte(config.Cstruct.AppServer.CACert),
		[]byte(config.Cstruct.AppServer.TLSCert), []byte(config.Cstruct.AppServer.TLSKey))
	if err != nil {
		return gateway, err
	}

	resp, err := appserverClient.GetGatewayByMac(context.Background(), &api.GetGatewayByMacRequest{Mac: mac})
	if err != nil {
		return gateway, err
	}

	walletId, err := db.Wallet.GetWalletIdFromOrgId(resp.OrgId)
	if err != nil {
		return gateway, err
	}

	createdTimeUpdated, _ := ptypes.Timestamp(resp.GwProfile.CreatedAt)
	gateway.Mac 		= resp.GwProfile.Mac
	gateway.FkWallet 	= walletId
	gateway.Mode 		= types.GW_WHOLE_NETWORK
	gateway.CreatedAt 	= createdTimeUpdated
	gateway.OrgId 		= resp.OrgId
	gateway.Description = resp.GwProfile.Description
	gateway.Name 		= resp.GwProfile.Name

	return gateway, nil
}

/*
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
}*/
