package supernode

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"strings"
)

var ethScan = connectEthScan()

func checkTokenTx(contractAddress, address, currAbv string) error {
	supernodeID, err := db.DbGetSuperNodeExtAccountId(config.Cstruct.SuperNode.ExtCurrAbv)

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

	for _, tx := range transfers {
		if strings.EqualFold(tx.To, address) && tx.BlockNumber > incurBlockNo {

			amount := float64(tx.Value.Int().Int64())
			_, err := db.DbAddTopUpRequest(tx.From, tx.To, tx.Hash, amount, currAbv)
			if err != nil {
				log.WithError(err).Warning("Storage: Cannot update TopUpRequest to DB")
				return err
			}

			// Update the last block to db
			err = db.DbUpdateLatestCheckedBlock(supernodeID, int64(tx.BlockNumber))
			if err != nil {
				log.WithError(err).Warning("Storage: Cannot update lastBlockNo to DB")
				return err
			}
			return nil
		}
	}
	return nil
}
