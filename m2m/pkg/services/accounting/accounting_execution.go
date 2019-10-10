package accounting

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func testAccounting(execTime time.Time, aggDurationMinutes int64, dlPrice float64) error {

	// aggStartAt := time.Now().UTC().AddDate(0, 0, -2) // change
	aggStartAt := execTime.Add(-time.Duration(aggDurationMinutes) * time.Minute)

	aggPeriodId, err := db.AggPeriod.InsertAggPeriod(aggStartAt, aggDurationMinutes, execTime)
	fmt.Println("InsertAggPeriod() ind: ", aggPeriodId, "  err:", err)

	MaxWalletId, errMaxWalletId := db.Wallet.GetMaxWalletId()
	if errMaxWalletId != nil {
		fmt.Println("GetMaxWalletId: ", MaxWalletId, "|| err:", errMaxWalletId) //@@
		return errMaxWalletId
	}

	awuList := make([]types.AggWltUsg, MaxWalletId+1)

	if err := getWltAggFromDlPkts(aggStartAt, aggDurationMinutes, awuList); err != nil {
		return err // @@ add path
	}

	addPricesWltAgg(awuList, dlPrice) // error does not matter

	addNonPriceFields(awuList, aggStartAt, aggDurationMinutes, aggPeriodId)

	walletIdSuperNode, errWltId := db.Wallet.GetWalletIdSuperNode()
	fmt.Println("*** walletIdSuperNode: ", walletIdSuperNode) //@@
	if errWltId != nil {
		log.Info("Error / unable to get superNodeAccount") // error
		return errWltId                                    // @@ wrap...
	}
	putInDbAggWltUsg(awuList, walletIdSuperNode)

	fmt.Printf("awuList: %+v\n", awuList)

	// forAggWltRecord
	// 	inserrtAggWlt
	return nil

}

func getWltAggFromDlPkts(aggStartAt time.Time, aggDurationMinutes int64, awuList []types.AggWltUsg) error {

	if wltIds, cnts, err := db.DlPacket.GetAggDlPktDeviceWallet(aggStartAt, aggDurationMinutes); true {
		fmt.Println("GetAggDlPktDeviceWallet   wltIds: ", wltIds, "  cnts: ", cnts, "   err: ", err)
		if err != nil {
			fmt.Println(err) // add path
			return err
		}
		if len(wltIds) != len(cnts) {
			fmt.Println("Inequal length of arrays wltIds, Cnts GetAggDlPktDeviceWallet") // @@ add path to error
		}
		for k, v := range wltIds {
			awuList[v].DlCntDv = cnts[k]
		}
	}

	if wltIds, cnts, err := db.DlPacket.GetAggDlPktGatewayWallet(aggStartAt, aggDurationMinutes); true {
		fmt.Println("GetAggDlPktGatewayWallet   wltIds: ", wltIds, "  cnts: ", cnts, "   err: ", err)
		if err != nil {
			fmt.Println(err) // add path
			//return   //@@
			return err
		}
		if len(wltIds) != len(cnts) {
			fmt.Println("Inequal length of arrays wltIds, Cnts GetAggDlPktGatewayWallet") // @@ add path to error
		}
		for k, v := range wltIds {
			awuList[v].DlCntGw = cnts[k]
		}
	}

	if wltIds, cnts, err := db.DlPacket.GetAggDlPktFreeWallet(aggStartAt, aggDurationMinutes); true {
		fmt.Println("GetAggDlPktFreeWallet   wltIds: ", wltIds, "  cnts: ", cnts, "   err: ", err)
		if err != nil {
			fmt.Println(err) // add path
			return err
		}
		if len(wltIds) != len(cnts) {
			fmt.Println("Inequal length of arrays wltIds, Cnts GetAggDlPktFreeWallet") // @@ add path to error
		}
		for k, v := range wltIds {
			awuList[v].DlCntDvFree = cnts[k]
			awuList[v].DlCntGwFree = cnts[k]
		}
	}
	return nil
}

func addPricesWltAgg(awuList []types.AggWltUsg, dlPrice float64) error {
	for k, v := range awuList {

		if v == (types.AggWltUsg{}) {
			continue
		}

		awuList[k].Spend = float64(v.DlCntDv-v.DlCntDvFree) * dlPrice
		awuList[k].Income = float64(v.DlCntGw-v.DlCntGwFree) * dlPrice
		awuList[k].BalanceIncrease = awuList[k].Income - awuList[k].Spend

	}
	return nil
}

func addNonPriceFields(awuList []types.AggWltUsg, aggStartAt time.Time, aggDurationMins int64, aggPeriodId int64) error {
	for k, v := range awuList {

		if v == (types.AggWltUsg{}) {
			continue
		}
		awuList[k].FkAggPeriod = aggPeriodId
		awuList[k].StartAt = aggStartAt
		awuList[k].DurationMinutes = aggDurationMins
		awuList[k].FkWallet = int64(k)
	}
	return nil
}

func putInDbAggWltUsg(awuList []types.AggWltUsg, walletIdSuperNode int64) error {
	fmt.Println("putIndDbAggWltUsg 1")

	for _, v := range awuList {

		if v == (types.AggWltUsg{}) {
			continue
		}

		insertedAggWltUsgId, errIns := db.AggWalletUsage.InsertAggWltUsg(v)
		if errIns != nil {
			fmt.Println("accounting/putInDbAggWltUsg impossible to write in DB ", errIns)
			/// return  //@@ to check return or not
		}

		fmt.Println("putIndDbAggWltUsg 2")
		_, err := db.AggWalletUsage.ExecAggWltUsgPayments(types.InternalTx{
			FkWalletSender: walletIdSuperNode,
			FkWalletRcvr:   v.FkWallet,
			PaymentCat:     string(types.DOWNLINK_AGGREGATION),
			TxInternalRef:  insertedAggWltUsgId,
			Value:          v.BalanceIncrease,
			TimeTx:         time.Now().UTC(),
		})
		if err != nil {
			fmt.Println("err ExecAggWltUsgPayments: ", err)
		}
	}
	return nil
}
