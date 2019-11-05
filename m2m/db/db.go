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

	// testDb()

	return nil
}

func dbInit() {
	dbErrorInit()

	if err := Wallet.CreateWalletTable(); err != nil {
		log.WithError(err).Fatal("db/CreateWalletTable")
	}

	if err := Wallet.CreateWalletFunctions(); err != nil {
		log.WithError(err).Fatal("db/CreateWalletFunctions")
	}

	if err := InternalTx.CreateInternalTxTable(); err != nil {
		log.WithError(err).Fatal("db/CreateInternalTxTable")
	}

	if err := ExtCurrency.CreateExtCurrencyTable(); err != nil {
		log.WithError(err).Fatal("db/CreateExtCurrencyTable")
	}

	if err := ExtAccount.CreateExtAccountTable(); err != nil {
		log.WithError(err).Fatal("db/CreateExtAccountTable")
	}

	if err := WithdrawFee.CreateWithdrawFeeTable(); err != nil {
		log.WithError(err).Fatal("db/CreateWithdrawFeeTable")
	}

	if err := Withdraw.CreateWithdrawTable(); err != nil {
		log.WithError(err).Fatal("db/CreateWithdrawTable")
	}

	if err := Withdraw.CreateWithdrawFunctions(); err != nil {
		log.WithError(err).Fatal("db/CreateWithdrawFunctions")
	}

	if err := Topup.CreateTopupTable(); err != nil {
		log.WithError(err).Fatal("db/CreateTopupTable")
	}

	if err := Topup.CreateTopupFunctions(); err != nil {
		log.WithError(err).Fatal("db/CreateTopupFunctions")
	}

	if err := Device.CreateDeviceTable(); err != nil {
		log.WithError(err).Fatal("db/CreateDeviceTable")
	}

	if err := Gateway.CreateGatewayTable(); err != nil {
		log.WithError(err).Fatal("db/CreateGatewayTable")
	}

	if err := AggPeriod.CreateAggPeriodTable(); err != nil {
		log.WithError(err).Fatal("db/CreateAggPeriodTable")
	}

	if err := AggWalletUsage.CreateAggWltUsgTable(); err != nil {
		log.WithError(err).Fatal("db/CreateAggWltUsgTable")
	}

	if err := AggWalletUsage.CreateAggWltUsgFunctions(); err != nil {
		log.WithError(err).Fatal("db/CreateAggWltUsgFunctions")
	}

	if err := AggDeviceUsage.CreateAggDvUsgTable(); err != nil {
		log.WithError(err).Fatal("db/CreateAggDvUsgTable")
	}

	if err := AggGatewayUsage.CreateAggGwUsgTable(); err != nil {
		log.WithError(err).Fatal("db/CreateAggGwUsgTable")
	}

	if err := DlPacket.CreateDlPktTable(); err != nil {
		log.WithError(err).Fatal("db/CreateDlPktTable")
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
