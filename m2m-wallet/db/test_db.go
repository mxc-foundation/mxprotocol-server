package db

import (
	"fmt"
	"time"

	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func testDb() {
	testWallet()
	// testInternalTx()
	// testWithdrawFee()
	// testExtCurrency()
	// testWithdraw()
	// testExtAccount()
	// testTopup()
}

func testWallet() {

	wi, errGetWI := DbGetWalletIdofActiveAcnt("0x7645", "MXC") //0x8347
	fmt.Println("DbGetWalletIdofActiveAcnt(): ", wi, " || err:", errGetWI)

	balance, err2 := DbGetWalletBalance(1)
	fmt.Println("GetWalletBalance(): ", balance, " || err:", err2)

	walletId, err := DbGetWalletIdFromOrgId(1)
	fmt.Println("GetWalletId(): ", walletId, " || err:", err)

	DbCreateWalletTable()

	retInd, errIns := DbInsertWallet(pstgDb.Wallet{
		FkOrgLa: 8,
		TypeW:   string(pstgDb.USER),
		Balance: 680.6})

	fmt.Println("DbInsertWallet  retInd:", retInd, "  err: ", errIns)

	var wp *pstgDb.Wallet = new(pstgDb.Wallet)
	DbGetWallet(wp, 2)
	fmt.Println("wp getWallet: ", *wp)

}

func testInternalTx() {
	DbCreateInternalTxTable()

	retInd, errIns := DbInsertInternalTx(
		pstgDb.InternalTx{
			FkWalletSender: 4,
			FkWalletRcvr:   21,
			PaymentCat:     string(pstgDb.TOP_UP),
			TxInternalRef:  4,
			Value:          65.23,
			TimeTx:         time.Now().UTC()})
	fmt.Println("err DbInsertInternalTx: ", retInd, "  err: ", errIns)

}

func testExtCurrency() {

	err := DbCreateExtCurrencyTable()
	fmt.Println("err DbCreateExtCurrencyTable(): ", err)

	_, err = DbInsertExtCurr(
		pstgDb.ExtCurrency{
			Name: "ethereum",
			Abv:  "ETH"})
	fmt.Println("err DbInsertExtCurr(): ", err)
}

func testWithdrawFee() {

	fee, errGetWF := DbGetActiveWithdrawFee("MXC")
	fmt.Println("DbGetActiveWithdrawFee(): ", fee, " err:", errGetWF)

	err := DbCreateWithdrawFeeTable()
	fmt.Println("err DbCreateWithdrawFeeTable(): ", err)

	wf := pstgDb.WithdrawFee{
		FkExtCurr:  1,
		Fee:        134.2,
		InsertTime: time.Now().UTC(),
		Status:     string(pstgDb.ACTIVE)}

	_, errIns := DbInsertWithdrawFee(wf)
	fmt.Println("err DbInsertWithdrawFee(): ", errIns)

}

func testExtAccount() {

	fmt.Println("DbUpdateLatestCheckedBlock(): err", DbUpdateLatestCheckedBlock(2, 876))

	blk, errBlk := DbGetLatestCheckedBlock(3)
	fmt.Println("DbGetLatestCheckedBlock(): ", blk, " err:", errBlk)

	acntAdr, errGetAu := DbGetUserExtAccountAdr(1, "MXC")
	fmt.Println("DbGetUserExtAccountAdr(): ", acntAdr, " err:", errGetAu)

	val, errGetAc := DbGetSuperNodeExtAccountAdr("MXC")
	fmt.Println("DbGetSuperNodeExtAccountAdr(): ", val, " err:", errGetAc)

	fmt.Println("err DbCreateExtAccountTable(): ", DbCreateExtAccountTable())

	ea := pstgDb.ExtAccount{
		FkWallet:           1,
		FkExtCurrency:      1,
		Account_adr:        "0x7645",
		Insert_time:        time.Now().UTC(),
		Status:             string(pstgDb.ARC),
		LatestCheckedBlock: 123}

	_, errIns := DBInsertExtAccount(ea)
	fmt.Println("err DBInsertExtAccount(): ", errIns)

}

func testWithdraw() {
	wdr := pstgDb.Withdraw{
		FkExtAcntSender:          1,
		FkExtAcntRcvr:            3,
		FkExtCurr:                1,
		Value:                    6.3,
		FkWithdrawFee:            1,
		TxSentTime:               time.Now().UTC(),
		TxStatus:                 string(pstgDb.PENDING),
		TxAprvdTime:              time.Now().UTC(),
		FkQueryIdePaymentService: 98,
		TxHash: "0x556664",
	}

	fmt.Println("err DbUpdateWithdrawSuccessful(): ", DbUpdateWithdrawSuccessful(11))

	fmt.Println("err DbCreateWithdrawSuccessfulFunction(): ", DbCreateWithdrawSuccessfulFunction())

	fmt.Println("err DbCreateWithdrawTable(): ", DbCreateWithdrawTable())

	_, errIns := DbInsertWithdraw(wdr)
	fmt.Println("err DbInsertWithdraw(): ", errIns)

	// it := pstgDb.InternalTx{
	// 	FkWalletSender: 1,
	// 	FkWalletRcvr:   35,
	// 	PaymentCat:     string(pstgDb.WITHDRAW),
	// 	TxInternalRef:  4,
	// 	Value:          65.23,
	// 	TimeTx:         time.Now().UTC()}

	// fmt.Println("err DbApplyWithdrawReq(): ", DbApplyWithdrawReq(wdr, it))
}

func testTopup() {
	fmt.Println("err DbCreateTopupTable(): ", DbCreateTopupTable())

	tu := pstgDb.Topup{
		FkExtAcntSender: 1,
		FkExtAcntRcvr:   2,
		FkExtCurr:       1,
		Value:           5.7,
		TxAprvdTime:     time.Now().UTC(),
		TxHash:          "0x456784",
	}
	retInd, err := DbInsertTopup(tu)
	if err == nil {
		tu.Id = retInd
	}

	fmt.Println("err DbInsertWithdraw(): x:", retInd, "err: ", err)

}
