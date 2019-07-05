package supernode

import (
	"fmt"
	"log"
	"strings"
)

var ethScan = connectEthScan()

func checkTokenTx(contractAddress, address string) {
	//ToDo: read lastTxNo from db
	currentBlockNo := 20

	transfers, err := ethScan.ERC20Transfers(&contractAddress, &address, &currentBlockNo, nil, 0, 0)
	if err != nil {
		log.Panic(err)
	}

	for _, tx := range transfers{
		if strings.EqualFold(tx.To, address) && tx.BlockNumber > currentBlockNo{
			fmt.Println("From:", tx.From)
			fmt.Println("TxHash: ", tx.Hash)
			fmt.Println("Amount: ", tx.Value.Int())
			fmt.Println("TimeStemp:", tx.TimeStamp.Time())
			//ToDo: rewrite the last block to db
			fmt.Println("BlockNo: ", tx.BlockNumber)
		}
	}
}