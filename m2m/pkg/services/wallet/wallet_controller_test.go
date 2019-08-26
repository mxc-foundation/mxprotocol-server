package wallet

import (
	"github.com/smartystreets/assertions"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/tests"
	"testing"
)

type WalletTestSuite services.ServicesTestSuite

func (s *WalletTestSuite) TESTGetWalletId(t *testing.T) {
	// setup database
	conf := tests.GetConfig()
	err := db.Setup(conf)
	assertions.ShouldBeNil(err)

	orgList := [10]int64{}
	for i, _ := range orgList {
		walletId, err := GetWalletId(int64(i))
		assertions.ShouldBeNil(err)

		orgList[i] = walletId
	}

	// new orgId, must return larger walletId
	for _, v := range orgList {
		walletId, err := GetWalletId(20)
		assertions.ShouldBeNil(err)
		assertions.ShouldBeGreaterThan(walletId, v)
	}

	// old orgId, must return same walletId
	for i, v := range orgList {
		walletId, err := GetWalletId(int64(i))
		assertions.ShouldBeNil(err)
		assertions.ShouldEqual(walletId, v)
	}

}
