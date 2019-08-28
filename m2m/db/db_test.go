package db

import (
	"github.com/stretchr/testify/suite"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/tests"
	"testing"
)

type DbInterfaceTestSuite struct {
	tx *TxHandler
}

// SetupSuite is called once before starting the test-suite.
func (b *DbInterfaceTestSuite) SetupSuite() {
	conf := tests.GetConfig()
	if err := Setup(conf); err != nil {
		panic(err)
	}
}

// SetupTest is called before every test.
func (b *DbInterfaceTestSuite) SetupTest() {
	tx, err := DB().Begin()
	if err != nil {
		panic(err)
	}
	b.tx = tx

	tests.ResetDB(DB().DB)
}

// TearDownTest is called after every test.
func (b *DbInterfaceTestSuite) TearDownTest() {
	if err := b.tx.Rollback(); err != nil {
		panic(err)
	}
}

// Tx returns a database transaction (which is rolled back after every
// test).
func (b *DbInterfaceTestSuite) Tx() *TxHandler {
	return b.tx
}

type DbTestSuite struct {
	suite.Suite
	DbInterfaceTestSuite
}

func TestDb(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

/*func testDb() {
	// testWallet()
	// testInternalTx()
	testWithdrawFee()
	// testExtCurrency()
	testWithdraw()
	// testExtAccount()
	testTopup()

}

func testWallet() {

	retInd, errIns := DbInsertWallet(0, SUPER_ADMIN) //@@ balance for the super node
	fmt.Println("DbInsertWallet  retInd:", retInd, "  err: ", errIns)

	retInd2, errIns2 := DbInsertWallet(3, USER) //@@ balance for the super node
	fmt.Println("DbInsertWallet  retInd:", retInd2, "  err: ", errIns2)

	superNodeId, errsuperNodeId := DbGetWalletIdSuperNode()
	fmt.Println("GetWalletIdSuperNode(): ", superNodeId, " || err:", errsuperNodeId)

	fmt.Println("err DbUpdateBalanceByWalletId(): ", DbUpdateBalanceByWalletId(1, 654.3))

	balance, err2 := DbGetWalletBalance(1)
	fmt.Println("GetWalletBalance(): ", balance, " || err:", err2)

	balance3, err3 := DbGetWalletBalance(2)
	fmt.Println("GetWalletBalance(): ", balance3, " || err:", err3)

	wi, errGetWI := DbGetWalletIdByActiveAcnt("0x7645", "MXC") //0x8347
	fmt.Println("DbGetWalletIdByActiveAcnt(): ", wi, " || err:", errGetWI)

	walletId, err := DbGetWalletIdFromOrgId(0)
	fmt.Println("GetWalletId(): ", walletId, " || err:", err)

	// var wp *pstgDb.Wallet = new(pstgDb.Wallet)
	// DbGetWallet(wp, 2)
	// fmt.Println("wp getWallet: ", *wp)

}

func testInternalTx() {
	dbCreateInternalTxTable()

	retInd, errIns := DbInsertInternalTx(
		InternalTx{
			FkWalletSender: 4,
			FkWalletRcvr:   21,
			PaymentCat:     string(TOP_UP),
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

	_, err := DbInsertExtCurr(
		ExtCurrency{
			Name: "MXC token",
			Abv:  "MXC"})
	fmt.Println("err DbInsertExtCurr(): ", err)

	idCur, errIdCur := DbGetExtCurrencyIdByAbbr("MXC")
	fmt.Println("DbGetExtCurrencyIdByAbbr(): ", idCur, " err:", errIdCur)

	err2 := dbCreateExtCurrencyTable()
	fmt.Println("err dbCreateExtCurrencyTable(): ", err2)

}

func testWithdrawFee() {

	err := dbCreateWithdrawFeeTable()
	fmt.Println("err dbCreateWithdrawFeeTable(): ", err)

	wf := WithdrawFee{
		FkExtCurr:  2,
		Fee:        9.99,
		InsertTime: time.Now().UTC(),
		Status:     string(ACTIVE)}

	_, errIns := DbInsertWithdrawFee("MXC", wf.Fee)
	fmt.Println("err DbInsertWithdrawFee(): ", errIns)

	// wid, errwid := DbGetActiveWithdrawFeeId("MXC")
	// fmt.Println("DbGetActiveWithdrawFeeId(): ", wid, " err:", errwid)

	fee, errGetWF := DbGetActiveWithdrawFee("MXC")
	fmt.Println("DbGetActiveWithdrawFee(): ", fee, " err:", errGetWF)

}

func testExtAccount() {

	indInsert, errIns := ext_account.DBInsertExtAccount(1, "0x11", "Ether")
	fmt.Println("err DBInsertExtAccount(): ", indInsert, errIns)

	acntId2, errGetAi2 := ext_account.DbGetExtAccountIdByAdr("0x1")
	fmt.Println("DbGetExtAccountIdByAdr(): ", acntId2, " err:", errGetAi2)
	// fmt.Println("suffix:", strings.HasSuffix(errGetAi2.Error(), DbError.NoRowQueryRes.Error()))

	acntId, errGetAi := ext_account.DbGetUserExtAccountId(2, "MXC")
	fmt.Println("DbGetUserExtAccountId(): ", acntId, " err:", errGetAi)
	acntAdr, errGetAu := ext_account.DbGetUserExtAccountAdr(1, "MXC")
	fmt.Println("DbGetUserExtAccountAdr(): ", acntAdr, " err:", errGetAu)

	valId, errGetids := ext_account.DbGetSuperNodeExtAccountId("MXC")
	fmt.Println("DbGetSuperNodeExtAccountId(): ", valId, " err:", errGetids)

	fmt.Println("DbUpdateLatestCheckedBlock(): err", ext_account.DbUpdateLatestCheckedBlock(3, 876))

	blk, errBlk := ext_account.DbGetLatestCheckedBlock(3)
	fmt.Println("DbGetLatestCheckedBlock(): ", blk, " err:", errBlk)

	val, errGetAc := ext_account.DbGetSuperNodeExtAccountAdr("MXC")
	fmt.Println("DbGetSuperNodeExtAccountAdr(): ", val, " err:", errGetAc)

	fmt.Println("err dbCreateExtAccountTable(): ", ext_account.dbCreateExtAccountTable())

	valHist, errHist := ext_account.DbGetExtAcntHist(2, 0, 100)
	fmt.Println("DbGetExtAcntHist() errHist: ", errHist)
	for _, v := range valHist {
		fmt.Println(v)
	}

}

func testWithdraw() {

	withId, errInitWith := DbInitWithdrawReq(2, 10, "MXC")
	fmt.Println(" DbInitWithdrawReq()  id: ", withId, "  err:", errInitWith)

	fmt.Println("err DbUpdateWithdrawPaymentQueryId(): ", DbUpdateWithdrawPaymentQueryId(withId, 111))

	fmt.Println("err DbUpdateWithdrawSuccessful(): ", DbUpdateWithdrawSuccessful(withId, "0x15525335", time.Now().UTC()))

	fmt.Println("err DbUpdateWithdrawPaymentQueryId(): ", DbUpdateWithdrawPaymentQueryId(10, 411))

	fmt.Println("err DbUpdateWithdrawSuccessful(): ", DbUpdateWithdrawSuccessful(10, "0x5435", time.Now().UTC()))

	valHist, errHist := DbGetWithdrawHist(2, 0, 100)
	fmt.Println("DbGetWithdrawHist() errHist: ", errHist)
	for _, v := range valHist {
		fmt.Println(v)
	}

	// fmt.Println("err dbCreateWithdrawTable(): ", dbCreateWithdrawTable())

	// _, errIns := DbInsertWithdraw(wdr)
	// fmt.Println("err DbInsertWithdraw(): ", errIns)

}

func testTopup() {

	topupId2, errAppl2 := DbAddTopUpRequest("0x222", "0x1", "0x231", 4, "MXC")
	fmt.Println("DbAddTopUpRequest() id: ", topupId2, "  err: ", errAppl2)

	topupId, errAppl := DbAddTopUpRequest("0x4", "0x1", "0x82363", 6, "MXC")
	fmt.Println("DbAddTopUpRequest() id: ", topupId, "  err: ", errAppl)

	valHist, errHist := DbGetTopupHist(4, 0, 100)
	fmt.Println("DbGetTopupHist() errHist: ", errHist)
	for _, v := range valHist {
		fmt.Println(v)
	}

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

}*/
