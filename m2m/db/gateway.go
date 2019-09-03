package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type GatewayMode string

const (
	GW_INACTIVE              GatewayMode = "INACTIVE"
	GW_FREE_GATEWAYS_LIMITED GatewayMode = "FREE_GATEWAYS_LIMITED"
	GW_WHOLE_NETWORK         GatewayMode = "WHOLE_NETWORK"
	GW_DELETED               GatewayMode = "DELETED"
)

type Gateway pg.Gateway

func DbCreateGatewayTable() error {
	return pg.PgDB.CreateGatewayTable()
}

func DbInsertGateway(mac string, fkGatewayNs int64, fkWallet int64, mode GatewayMode, orgId int64, description string, name string) (insertIndex int64, err error) {
	gw := pg.Gateway{
		Mac:         mac,
		FkGatewayNs: fkGatewayNs,
		FkWallet:    fkWallet,
		Mode:        string(mode),
		CreatedAt:   time.Now().UTC(),
		OrgId:       orgId,
		Description: description,
		Name:        name,
	}
	return pg.PgDB.InsertGateway(gw)
}

func DbGetGatewayMode(gwId int64) (gwMode GatewayMode, err error) {
	return GW_INACTIVE, nil
}

func DbSetGatewayMode(gwId int64, gwMode GatewayMode) (err error) {
	return nil
}

func DbDeletGateway(gwId int64) (err error) {
	return DbSetGatewayMode(gwId, GW_DELETED)
}

func DbGetGatewayIdByMac(mac string) (gwId int64, err error) {
	return 1, nil
}

func DbUpdateGatewayLastSeen(gwId int64, newTime time.Time) (err error) {
	return nil
}

func DbGetGatewayProfile(gwId int64) (gw Gateway, err error) {
	return Gateway{}, nil
}

func castGatewayList(gwList []pg.Gateway, err error) ([]Gateway, error) {
	var castedVal []Gateway
	for _, v := range gwList {
		castedVal = append(castedVal, Gateway(v))
	}
	return castedVal, err
}

func DbGetGatewayListByWalletId(walletId int64) (gwList []Gateway, err error) {
	return nil, nil
}

func DbGetGatewayRecCnt(walletId int64) (recCnt int64, err error) {
	// return pg.PgDB.GetDeviceRecCnt(walletId)
	return 1, nil
}
