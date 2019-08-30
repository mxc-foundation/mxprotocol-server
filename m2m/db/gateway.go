package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type GatewayeMode string

const (
	GW_INACTIVE              GatewayeMode = "INACTIVE"
	GW_FREE_GATEWAYS_LIMITED GatewayeMode = "FREE_GATEWAYS_LIMITED"
	GW_WHOLE_NETWORK         GatewayeMode = "WHOLE_NETWORK"
	GW_DELETED               GatewayeMode = "DELETED"
)

func DbCreateGatewayTable() error {
	return pg.PgDB.CreateGatewayTable()
}

func DbInsertGateway(mac string, fkGatewayNs int64, fkWallet int64, mode DeviceMode, orgId int64, description string, name string) (insertIndex int64, err error) {
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
