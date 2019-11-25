package db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	var InitDbTableTasks = []func() error{

		ConfigTable.CreateConfigTable,
		Wallet.CreateWalletTable,
		Wallet.CreateWalletFunctions,
		InternalTx.CreateInternalTxTable,
		ExtCurrency.CreateExtCurrencyTable,
		ExtAccount.CreateExtAccountTable,
		WithdrawFee.CreateWithdrawFeeTable,
		Withdraw.CreateWithdrawTable,
		Withdraw.CreateWithdrawFunctions,
		Topup.CreateTopupTable,
		Topup.CreateTopupFunctions,
		Device.CreateDeviceTable,
		Gateway.CreateGatewayTable,
		DlPacket.CreateDlPktTable,
		AggPeriod.CreateAggPeriodTable,
		AggWalletUsage.CreateAggWltUsgTable,
		AggWalletUsage.CreateAggWltUsgFunctions,
		AggDeviceUsage.CreateAggDvUsgTable,
		AggGatewayUsage.CreateAggGwUsgTable,
		Stake.CreateStakeTable,
		Stake.CreateStakeFunctions,
		StakeRevenuePeriod.CreateStakeRevenuePeriodTable,
		StakeRevenue.CreateStakeRevenueTable,
		StakeRevenue.CreateStakeRevenueFunctions,
	}

	for _, t := range InitDbTableTasks {
		if err := t(); err != nil {
			log.WithError(err).Fatal()
		}
	}

	if err := ConfigTable.CreateConfigTable(); err != nil {
		log.WithError(err).Fatal("db/CreateConfigTable")
	} else {
		// set default config

		config := map[string]interface{}{
			"downlink_fee":                 viper.GetInt("pricing.downlink_fee"),
			"transaction_percentage_share": viper.GetInt("pricing.transaction_percentage_share"),
			"low_balance_warning":          viper.GetInt("system_notification.low_balance_warning"),
		}

		if err = ConfigTable.InsertConfigs(config, true); err != nil {
			log.WithError(err).Fatal("db/InsertConfigs")
		}
	}
}
