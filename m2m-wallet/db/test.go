package db

import (
	"fmt"

	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func testDb() {
	testWallet()
}

func testWallet() {

	fmt.Println("GetWalletId(1): ", DbGetWalletId(1))

	DbCreateWalletTable()

	DbInsertWallet(pstgDb.Wallet{
		FkOrgLa: 3,
		TypeW:   pstgDb.USER,
		Balance: 100})

	var wp *pstgDb.Wallet = new(pstgDb.Wallet)
	DbGetWallet(wp, 1)
	fmt.Println("wp getWallet: ", *wp)

}
