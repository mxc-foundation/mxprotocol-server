package supernode

import (
	"fmt"
	"log"
	"strings"
)

var ethScan = connectEthScan()
//var client = connectMainClient()

func checkTokenTx(contractAddress, address string) {
	//ToDo: read lastTxNo from db
	currentBlockNo := 20

	/*block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}*/

	// check ERC20 transactions from/to a specified address
	transfers, err := ethScan.ERC20Transfers(&contractAddress, &address, &currentBlockNo, nil, 0, 0)
	if err != nil {
		log.Panic(err)
	}

	for i, tx := range transfers{
		if strings.EqualFold(tx.To, "0x8a96E17d85Bd897a88B547718865de990D2Fcb80") && tx.BlockNumber > currentBlockNo{
			fmt.Println("Tx: ", i+1)





			fmt.Println("Who :", tx.From)
			fmt.Println("TxHash: ", tx.Hash)
			fmt.Println("Amount: ", tx.Value.Int())
			fmt.Println("TimeStemp:", tx.TimeStamp.Time())
			//ToDo: rewrite the last block to db
			fmt.Println("BlockNo: ", tx.BlockNumber)
		}
	}
}