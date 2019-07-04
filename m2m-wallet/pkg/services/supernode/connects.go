package supernode

import (
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
)

/*func connectMainClient() *ethclient.Client {

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	_ = client //??
	return client
}*/

func connectEthScan() *etherscan.Client {
	tokenEthScan := etherscan.New(etherscan.Mainnet, "W8M6B92HBM7CUAQINJ8IMST29RY2ZVSQH4")

	return tokenEthScan
}