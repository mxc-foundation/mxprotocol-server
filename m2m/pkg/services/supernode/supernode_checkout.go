package supernode

import (
	"github.com/nanmu42/etherscan-api"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"math"
	"math/big"
	"strings"
)

func checkTokenTx(contractAddress, address, currAbv string) error {
	var ethScan *etherscan.Client
	etherTestNet := config.Cstruct.SuperNode.TestNet

	if etherTestNet == true {
		ethScan = connectEthTestScan()
	} else {
		ethScan = connectEthScan()
	}

	supernodeID, err := db.DbGetSuperNodeExtAccountId(config.Cstruct.SuperNode.ExtCurrAbv)
	if err != nil {
		log.WithError(err).Warning("storage: Cannot get supernodeID from DB")
		return err
	}

	currentBlockNo, err := db.DbGetLatestCheckedBlock(supernodeID)
	if err != nil {
		log.WithError(err).Warning("storage: Cannot get currentBlockNo from DB")
		return err
	}

	incurBlockNo := int(currentBlockNo)

	transfers, err := ethScan.ERC20Transfers(&contractAddress, &address, &incurBlockNo, nil, 0, 0)
	if err != nil {
		log.WithError(err).Warning("Etherscan: Cannot get reply from Etherscan")
		return err
	}

	var newBlockNo int64

	for _, tx := range transfers {
		if strings.EqualFold(tx.To, address) && tx.BlockNumber > incurBlockNo {
			fbalance := new(big.Float)
			fbalance.SetString(tx.Value.Int().String())
			ethValue, _ := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18))).Float64()

			from, err := db.DbGetExtAccountIdByAdr(tx.From, config.Cstruct.SuperNode.ExtCurrAbv) //config.Cstruct.SuperNode.ExtCurrAbv is added by Aslan => please check!
			if err != nil {
				log.WithError(err).Warning("Cannot get external account from DB")
				return err
			}
			if from == 0 {
				log.WithError(err).Warning(tx.From, " is not in DB")
				continue
			}

			to, err := db.DbGetExtAccountIdByAdr(tx.To, config.Cstruct.SuperNode.ExtCurrAbv) //config.Cstruct.SuperNode.ExtCurrAbv is added by Aslan => please check!
			if err != nil {
				log.WithError(err).Warning("Cannot get super node account from DB")
				return err
			}
			if to == 0 {
				log.WithError(err).Warning(tx.To, " is not in DB")
				continue
			}

			_, err = db.DbAddTopUpRequest(tx.From, tx.To, tx.Hash, ethValue, currAbv)
			if err != nil {
				log.WithError(err).Warning("Storage: Cannot update TopUpRequest to DB")
				return err
			}
		}
		newBlockNo = int64(tx.BlockNumber)
	}

	// Update the last block to db
	if newBlockNo > currentBlockNo {
		err = db.DbUpdateLatestCheckedBlock(supernodeID, int64(newBlockNo))
		if err != nil {
			log.WithError(err).Warning("Storage: Cannot update lastBlockNo to DB")
			return err
		}
	}

	return nil
}
