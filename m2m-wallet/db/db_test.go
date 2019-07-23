package db

import (
	"fmt"
	"strings"
	"time"

	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func testDb() {
	// testWallet()
	// testInternalTx()
	testWithdrawFee()
	// testExtCurrency()
	// testWithdraw()
	// testExtAccount()
	// testTopup()

}

func testWallet() {

	superNodeId, errsuperNodeId := DbGetWalletIdSuperNode()
	fmt.Println("GetWalletIdSuperNode(): ", superNodeId, " || err:", errsuperNodeId)

	fmt.Println("err DbUpdateBalanceByWalletId(): ", DbUpdateBalanceByWalletId(1, 654.3))

	balance, err2 := DbGetWalletBalance(102)
	fmt.Println("GetWalletBalance(): ", balance, " || err:", err2)

	wi, errGetWI := DbGetWalletIdByActiveAcnt("0x7645", "MXC") //0x8347
	fmt.Println("DbGetWalletIdByActiveAcnt(): ", wi, " || err:", errGetWI)

	walletId, err := DbGetWalletIdFromOrgId(1)
	fmt.Println("GetWalletId(): ", walletId, " || err:", err)

	dbCreateWalletTable()

	retInd, errIns := DbInsertWallet(10, USER) //@@ balance for the super node

	fmt.Println("DbInsertWallet  retInd:", retInd, "  err: ", errIns)

	// var wp *pstgDb.Wallet = new(pstgDb.Wallet)
	// DbGetWallet(wp, 2)
	// fmt.Println("wp getWallet: ", *wp)

}

func testInternalTx() {
	dbCreateInternalTxTable()

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

	// _, err = DbInsertExtCurr(
	// 	pstgDb.ExtCurrency{
	// 		Name: "ethereum",
	// 		Abv:  "ETH"})
	// fmt.Println("err DbInsertExtCurr(): ", err)

	idCur, errIdCur := DbGetExtCurrencyIdByAbbr("MXC")
	fmt.Println("DbGetExtCurrencyIdByAbbr(): ", idCur, " err:", errIdCur)

	err := dbCreateExtCurrencyTable()
	fmt.Println("err dbCreateExtCurrencyTable(): ", err)

	_, err = DbInsertExtCurr(
		pstgDb.ExtCurrency{
			Name: "ethereum",
			Abv:  "ETH"})
	fmt.Println("err DbInsertExtCurr(): ", err)
}

func testWithdrawFee() {

	// wid, errwid := DbGetActiveWithdrawFeeId("MXC")
	// fmt.Println("DbGetActiveWithdrawFeeId(): ", wid, " err:", errwid)

	// fee, errGetWF := DbGetActiveWithdrawFee("MXC")
	// fmt.Println("DbGetActiveWithdrawFee(): ", fee, " err:", errGetWF)

	err := dbCreateWithdrawFeeTable()
	fmt.Println("err dbCreateWithdrawFeeTable(): ", err)

	wf := pstgDb.WithdrawFee{
		FkExtCurr:  1,
		Fee:        99.1,
		InsertTime: time.Now().UTC(),
		Status:     string(pstgDb.ACTIVE)}

	_, errIns := DbInsertWithdrawFee("MXC", wf.Fee)
	fmt.Println("err DbInsertWithdrawFee(): ", errIns)

}

func testExtAccount() {

	ea := pstgDb.ExtAccount{
		FkWallet:           1,
		FkExtCurrency:      1,
		Account_adr:        "0x616",
		Insert_time:        time.Now().UTC(),
		Status:             string(pstgDb.ARC),
		LatestCheckedBlock: 123}

	indInsert, errIns := DBInsertExtAccount(int64(ea.FkWallet), ea.Account_adr, "Ether")
	fmt.Println("err DBInsertExtAccount(): ", indInsert, errIns)

	acntId2, errGetAi2 := DbGetExtAccountIdByAdr("0x8347")
	fmt.Println("DbGetExtAccountIdByAdr(): ", acntId2, " err:", errGetAi2)
	fmt.Println("suffix:", strings.HasSuffix(errGetAi2.Error(), DbError.NoRowQueryRes.Error()))

	acntId, errGetAi := DbGetUserExtAccountId(1, "MXC")
	fmt.Println("DbGetUserExtAccountId(): ", acntId, " err:", errGetAi)
	acntAdr, errGetAu := DbGetUserExtAccountAdr(1, "MXC")
	fmt.Println("DbGetUserExtAccountAdr(): ", acntAdr, " err:", errGetAu)

	valId, errGetids := DbGetSuperNodeExtAccountId("MXC")
	fmt.Println("DbGetSuperNodeExtAccountId(): ", valId, " err:", errGetids)

	fmt.Println("DbUpdateLatestCheckedBlock(): err", DbUpdateLatestCheckedBlock(2, 876))

	blk, errBlk := DbGetLatestCheckedBlock(3)
	fmt.Println("DbGetLatestCheckedBlock(): ", blk, " err:", errBlk)

	val, errGetAc := DbGetSuperNodeExtAccountAdr("MXC")
	fmt.Println("DbGetSuperNodeExtAccountAdr(): ", val, " err:", errGetAc)

	fmt.Println("err dbCreateExtAccountTable(): ", dbCreateExtAccountTable())

}

func testWithdraw() {

	withId, errInitWith := DbInitWithdrawReq(1, 99, "MXC")
	fmt.Println(" DbInitWithdrawReq()  id: ", withId, "  err:", errInitWith)

	fmt.Println("err DbUpdateWithdrawPaymentQueryId(): ", DbUpdateWithdrawPaymentQueryId(1, 111))

	// wdr := pstgDb.Withdraw{
	// 	FkExtAcntSender:          1,
	// 	FkExtAcntRcvr:            2,
	// 	FkExtCurr:                1,
	// 	Value:                    45.2,
	// 	FkWithdrawFee:            1,
	// 	TxSentTime:               time.Now().UTC(),
	// 	TxStatus:                 string(pstgDb.PENDING),
	// 	TxAprvdTime:              time.Now().UTC(),
	// 	FkQueryIdePaymentService: 6,
	// 	TxHash: "0x06666344",
	// }

	fmt.Println("err DbUpdateWithdrawSuccessful(): ", DbUpdateWithdrawSuccessful(11, "0x555335", time.Now().UTC()))

	// fmt.Println("err dbCreateWithdrawTable(): ", dbCreateWithdrawTable())

	// _, errIns := DbInsertWithdraw(wdr)
	// fmt.Println("err DbInsertWithdraw(): ", errIns)

	// 	it := pstgDb.InternalTx{
	// 		FkWalletSender: 1,
	// 		FkWalletRcvr:   2,
	// 		PaymentCat:     string(pstgDb.WITHDRAW),
	// 		TxInternalRef:  4,
	// 		Value:          65.23,
	// 		TimeTx:         time.Now().UTC()}

	// 	fmt.Println("err DbApplyWithdrawReq(): ", DbInitWithdrawReqApply(wdr, it))
}

func testTopup() {

	topupId2, errAppl2 := DbAddTopUpRequest("0x8347", "0x7645", "0x5200006", 33.13, "MXC")
	fmt.Println("DbAddTopUpRequest() id: ", topupId2, "  err: ", errAppl2)

	// tu := pstgDb.Topup{
	// 	FkExtAcntSender: 2,
	// 	FkExtAcntRcvr:   1,
	// 	FkExtCurr:       1,
	// 	Value:           99.7,
	// 	TxAprvdTime:     time.Now().UTC(),
	// 	TxHash:          "0x78651111",
	// }

	// it := pstgDb.InternalTx{
	// 	FkWalletSender: 1,
	// 	FkWalletRcvr:   2,
	// 	PaymentCat:     string(pstgDb.TOP_UP),
	// }

	// topupId, errAppl := DbApplyTopup(tu, it)
	// fmt.Println("DbApplyTopup() id: ", topupId, "  err: ", errAppl)

	// retInd, err := DbInsertTopup(tu)
	// if err == nil {
	// 	tu.Id = retInd
	// }

	// fmt.Println("err dbCreateTopupTable(): ", dbCreateTopupTable())
	// fmt.Println("err DbInsertWithdraw(): x:", retInd, "err: ", err)

}
