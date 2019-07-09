package supernode

import (
	"fmt"
	"log"
	"strings"
)

var ethScan = connectEthScan()

func checkTokenTx(contractAddress, address string) {
	//ToDo: wait (get supernod ID)
	//supernodeID, err := db.DbGetSuperNodeExtAccountId(config.Cstruct.SuperNode.ExtCurrAbv)

	//currentBlockNo, err := db.DbGetLatestCheckedBlock(supernodeID)

	//For test!
	currentBlockNo := 20

	transfers, err := ethScan.ERC20Transfers(&contractAddress, &address, &currentBlockNo, nil, 0, 0)
	if err != nil {
		log.Panic(err)
	}

	for _, tx := range transfers {
		if strings.EqualFold(tx.To, address) && tx.BlockNumber > currentBlockNo {
			fmt.Println("From:", tx.From)
			fmt.Println("TxHash: ", tx.Hash)
			fmt.Println("Amount: ", tx.Value.Int())
			fmt.Println("TimeStemp:", tx.TimeStamp.Time())
			fmt.Println("BlockNo: ", tx.BlockNumber)

			//ToDo: wait (update the last block to db)
			//db.DbAddTopUpRequest(tx.From,tx.To,tx.Hash,tx.Value)

			//db.DbUpdateLatestCheckedBlock(supernodeID, tx.BlockNumber)
		}
	}
}
