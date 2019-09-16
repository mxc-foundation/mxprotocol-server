package supernode

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/ext_account"
)

func Setup(conf config.MxpConfig) error {
	supernode_account, err := ext_account.GetActiveExtAccount(0, conf.SuperNode.ExtCurrAbv)
	if err != nil {
		return err
	} else if supernode_account == "" {
		err := ext_account.UpdateActiveExtAccount(0, conf.SuperNode.SuperNodeAddress, conf.SuperNode.ExtCurrAbv)
		if err != nil {
			return err
		}
	}

	ticker_superAccount := time.NewTicker(time.Duration(config.Cstruct.SuperNode.CheckAccountSeconds) * time.Second)
	go func() {
		log.Info("Start supernode goroutine")
		for range ticker_superAccount.C {
			supernodeAccount, err := ext_account.GetActiveExtAccount(0, conf.SuperNode.ExtCurrAbv)
			if err != nil {
				log.WithError(err).Warning("service/supernode")
				continue
			}

			err = checkTokenTx(conf.SuperNode.ContractAddress, supernodeAccount, conf.SuperNode.ExtCurrAbv)
			if err != nil {
				log.Warning("Restarting...")
				continue
			}
		}
	}()

	log.Info("setup supernode service")
	return nil
}

