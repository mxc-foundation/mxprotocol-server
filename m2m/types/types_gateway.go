package types

import "time"

type Gateway struct {
	Id          int64       `db:"id"`
	Mac         string      `db:"mac"` // fk in AS (App Server)
	FkGatewayNs int64       `db:"fk_gateway_ns"`
	FkWallet    int64       `db:"fk_wallet"`
	Mode        GatewayMode `db:"mode"`
	CreatedAt   time.Time   `db:"created_at"`
	LastSeenAt  time.Time   `db:"last_seen_at"`
	OrgId       int64       `db:"org_id"`
	Description string      `db:"description"`
	Name        string      `db:"name"`
}

type GatewayMode string

const (
	GW_INACTIVE              GatewayMode = "INACTIVE"
	GW_FREE_GATEWAYS_LIMITED GatewayMode = "FREE_GATEWAYS_LIMITED"
	GW_WHOLE_NETWORK         GatewayMode = "WHOLE_NETWORK"
	GW_DELETED               GatewayMode = "DELETED"
)
