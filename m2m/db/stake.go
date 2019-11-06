package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type stakeDBInterface interface {
	CreateStakeTable() error
	InsertStake(walletId int64, amount float64) (insertIndex int64, err error)
	Unstake(stakeId int64) error
	GetActiveStake(walletId int64) (stakeProfile types.Stake, err error)
	GetStakeHistory(walletId int64, offset int64, limit int64) (stakeProfiles []types.Stake, err error)
	GetStakeHistoryCnt(walletId int64) (recCnt int64, err error)
	GetActiveStakes() (stakeProfiles []types.Stake, err error)
}

var Stake = stakeDBInterface(&pg.PgStake)
