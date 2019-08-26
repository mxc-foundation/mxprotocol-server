package supernode

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"log"
)

func connectMainClient() *ethclient.Client {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	_ = client
	return client
}

func connectEthTestScan() *etherscan.Client {
	tokenEthScan := etherscan.New(etherscan.Ropsten, config.Cstruct.SuperNode.APIKey)

	return tokenEthScan
}

func connectEthScan() *etherscan.Client {
	tokenEthScan := etherscan.New(etherscan.Mainnet, config.Cstruct.SuperNode.APIKey)

	return tokenEthScan
}
