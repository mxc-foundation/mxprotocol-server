package db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/migrations"
)

func Setup(conf config.MxpConfig) error {

	i = &dbM2M
	err := openDB(i, conf)
	if err != nil {
		return err
	}

	dbInit()

	if conf.PostgreSQL.Automigrate {
		log.Info("db/applying PostgreSQL data migrations")
		m := &migrate.AssetMigrationSource{
			Asset:    migrations.Asset,
			AssetDir: migrations.AssetDir,
			Dir:      "",
		}
		n, err := migrate.Exec(dbM2M.DB, "postgres", m, migrate.Up)
		if err != nil {
			return errors.Wrap(err, "db/applying PostgreSQL data migrations error")
		}
		log.WithField("count", n).Info("db/PostgreSQL data migrations applied")
	}

	return nil
}

func dbInit() {
	dbErrorInit()

	if err := ConfigTable.CreateConfigTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := Wallet.CreateWalletTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := Wallet.CreateWalletFunctions(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := InternalTx.CreateInternalTxTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := ExtCurrency.CreateExtCurrencyTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := ExtAccount.CreateExtAccountTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := WithdrawFee.CreateWithdrawFeeTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := Withdraw.CreateWithdrawTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := Withdraw.CreateWithdrawFunctions(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := Topup.CreateTopupTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := Topup.CreateTopupFunctions(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := Device.CreateDeviceTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := Gateway.CreateGatewayTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := DlPacket.CreateDlPktTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := AggPeriod.CreateAggPeriodTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := AggWalletUsage.CreateAggWltUsgTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := AggWalletUsage.CreateAggWltUsgFunctions(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := AggDeviceUsage.CreateAggDvUsgTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := AggGatewayUsage.CreateAggGwUsgTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := Stake.CreateStakeTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := StakeRevenuePeriod.CreateStakeRevenuePeriodTable(); err != nil {
		log.WithError(err).Fatal()
	}

	if err := StakeRevenue.CreateStakeRevenueTable(); err != nil {
		log.WithError(err).Fatal()
	}

}
