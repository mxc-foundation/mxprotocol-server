package db

import (
	"fmt"
	"time"

	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func testDb() {
	// testWallet()
	// testInternalTx()
	// testExtCurrency()
	// testWithdrawFee()
	// testWithdraw()
	testExtAccount()
}

func testWallet() {

	balance, err2 := DbGetWalletBalance(1)
	fmt.Println("GetWalletBalance(): ", balance, " || err:", err2)

	walletId, err := DbGetWalletIdFromOrgId(1)
	fmt.Println("GetWalletId(): ", walletId, " || err:", err)

	DbCreateWalletTable()

	DbInsertWallet(pstgDb.Wallet{
		FkOrgLa: 5,
		TypeW:   string(pstgDb.USER),
		Balance: 700})

	var wp *pstgDb.Wallet = new(pstgDb.Wallet)
	DbGetWallet(wp, 2)
	fmt.Println("wp getWallet: ", *wp)

}

func testInternalTx() {
	DbCreateInternalTxTable()

	fmt.Println("err DbInsertInternalTx: ",
		DbInsertInternalTx(
			pstgDb.InternalTx{
				FkWalletSender: 1,
				FkWalletRcvr:   2,
				PaymentCat:     string(pstgDb.WITHDRAW),
				TxInternalRef:  2,
				Value:          990.4,
				TimeTx:         time.Now().UTC()}))

}

func testExtCurrency() {

	err := DbCreateExtCurrencyTable()
	fmt.Println("err DbCreateExtCurrencyTable(): ", err)

	err = DbInsertExtCurr(
		pstgDb.ExtCurrency{
			Name: "MXC_token",
			Abv:  "MXC"})
	fmt.Println("err DbInsertExtCurr(): ", err)
}

func testWithdrawFee() {
	err := DbCreateWithdrawFeeTable()
	fmt.Println("err DbCreateWithdrawFeeTable(): ", err)

	wf := pstgDb.WithdrawFee{
		FkExtCurr:  1,
		Fee:        34.2,
		InsertTime: time.Now().UTC(),
		Status:     "ACTIVE"}

	fmt.Println("err DbInsertExtCurr(): ", DbInsertWithdrawFee(&wf))

}

func testExtAccount() {
	fmt.Println("err DbCreateExtAccountTable(): ", DbCreateExtAccountTable())

	ea := pstgDb.ExtAccount{
		FkWallet:      2,
		FkExtCurrency: 1,
		Account_adr:   "0x1234",
		Insert_time:   time.Now().UTC(),
		Status:        "ACTIVE"}

	fmt.Println("err DBInsertExtAccount(): ", DBInsertExtAccount(&ea))

}

func testWithdraw() {
	fmt.Println("err DbCreateWithdrawTable(): ", DbCreateWithdrawTable())

}
