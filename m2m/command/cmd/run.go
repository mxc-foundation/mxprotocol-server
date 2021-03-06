package cmd

import (
	"context"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/api/clients/appserver"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/staking"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/api"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/config"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/accounting"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/device"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/ext_account"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/gateway"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/supernode"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/topup"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/wallet"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/withdraw"
)

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tasks := []func() error{
		setLogLevel,
		printStartMessage,
		setupDb,
		setupAppserver,
		setupMoney,
		setupWallet,
		setupWithdraw,
		setupTopUp,
		setupSupernode,
		setupAPI,
		setupDevice,
		setupGateway,
		setupAccounting,
		setupStaking,
	}

	for _, t := range tasks {
		if err := t(); err != nil {
			log.Fatal(err)
		}
	}

	sigChan := make(chan os.Signal)
	exitChan := make(chan struct{})
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.WithField("signal", <-sigChan).Info("signal received")
	go func() {
		log.Warning("stopping mxp-server")
		exitChan <- struct{}{}
	}()
	select {
	case <-exitChan:
	case s := <-sigChan:
		log.WithField("signal", s).Info("signal received, stopping immediately")
	}

	return nil
}

func setLogLevel() error {
	log.SetLevel(log.Level(uint8(config.Cstruct.General.LogLevel)))
	return nil
}

func printStartMessage() error {
	log.WithFields(log.Fields{
		"version": version,
	}).Info("starting mxp-server")
	return nil
}

func setupDb() error {
	if err := db.Setup(config.Cstruct); err != nil {
		return errors.Wrap(err, "setup db error")
	}
	return nil
}

func setupAppserver() error {
	if err := appserver.Setup(config.Cstruct); err != nil {
		return errors.Wrap(err, "setup appserver client pool error")
	}
	return nil
}

func setupWithdraw() error {
	if err := withdraw.Setup(config.Cstruct); err != nil {
		return errors.Wrap(err, "setup service withdraw error")
	}
	return nil
}

func setupWallet() error {
	if err := wallet.Setup(config.Cstruct); err != nil {
		return errors.Wrap(err, "setup service wallet error")
	}
	return nil
}

func setupTopUp() error {
	if err := topup.Setup(); err != nil {
		return errors.Wrap(err, "setup service top_up error")
	}
	return nil
}

func setupSupernode() error {
	if err := supernode.Setup(config.Cstruct); err != nil {
		return errors.Wrap(err, "setup service super_node error")
	}
	return nil
}

func setupMoney() error {
	if err := ext_account.Setup(); err != nil {
		return errors.Wrap(err, "setup service ext_account error")
	}
	return nil
}

func setupDevice() error {
	if err := device.Setup(); err != nil {
		return errors.Wrap(err, "setup service device error")
	}
	return nil
}

func setupGateway() error {
	if err := gateway.Setup(); err != nil {
		return errors.Wrap(err, "setup service gateway error")
	}
	return nil
}

func setupAPI() error {
	if err := api.SetupHTTPServer(config.Cstruct); err != nil {
		return errors.Wrap(err, "setup api error")
	}
	return nil
}

func setupAccounting() error {
	if err := accounting.Setup(config.Cstruct); err != nil {
		return errors.Wrap(err, "setup service accounting error")
	}
	return nil
}

func setupStaking() error {
	if err := staking.Setup(config.Cstruct); err != nil {
		return errors.Wrap(err, "setup service staking error")
	}
	return nil
}
