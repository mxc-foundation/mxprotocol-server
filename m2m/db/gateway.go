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

func DbGetGatewayMode(gwId int64) (gwMode string, err error) {
	return pg.PgDB.GetGatewayMode(gwId)
}

func DbSetGatewayMode(gwId int64, gwMode GatewayMode) (err error) {
	return pg.PgDB.SetGatewayMode(gwId, string(gwMode))
}

func DbDeleteGateway(gwId int64) (err error) {
	return DbSetGatewayMode(gwId, GW_DELETED)
}

func DbGetGatewayIdByMac(mac string) (gwId int64, err error) {
	return pg.PgDB.GetGatewayIdByMac(mac)
}

func DbUpdateGatewayLastSeen(gwId int64, newTime time.Time) (err error) {
	return pg.PgDB.UpdateGatewayLastSeen(gwId, newTime)
}

func DbGetGatewayProfile(gwId int64) (gw Gateway, err error) {
	prf, err := pg.PgDB.GetGatewayProfile(gwId)
	return Gateway(prf), err
}

func castGatewayList(gwList []pg.Gateway, err error) ([]Gateway, error) {
	var castedVal []Gateway
	for _, v := range gwList {
		castedVal = append(castedVal, Gateway(v))
	}
	return castedVal, err
}

func DbGetGatewayListOfWallet(walletId int64, offset int64, limit int64) (gwList []Gateway, err error) {
	return castGatewayList(pg.PgDB.GetGatewayListOfWallet(walletId, offset, limit))
}

func DbGetGatewayRecCnt(walletId int64) (recCnt int64, err error) {
	return pg.PgDB.GetGatewayRecCnt(walletId)
}
