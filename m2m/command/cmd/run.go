package cmd

import (
	"context"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/device"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/gateway"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/ext_account"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/supernode"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/topup"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/wallet"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/withdraw"
)

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tasks := []func() error{
		setLogLevel,
		printStartMessage,
		setupDb,
		setupAuth,
		setupMoney,
		setupWallet,
		setupWithdraw,
		setupTopUp,
		setupSupernode,
		setupDevice,
		setupGateway,
		setupAPI,
		startM2MAPIServer,
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

func setupAuth() error {
	if err := auth.Setup(config.Cstruct); err != nil {
		return errors.Wrap(err, "Setup auth error")
	}
	return nil
}

func setupDb() error {
	if err := db.Setup(config.Cstruct); err != nil {
		return errors.Wrap(err, "setup db error")
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