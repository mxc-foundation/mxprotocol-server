package accounting

import (
	"fmt"
	"time"

	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func testAccounting() error {

	aggStartAt := time.Now().UTC().AddDate(0, 0, -1) // change
	aggDurationMinutes := int64(48 * 60)
	var aggPeriodId int64
	var dlPrice float64 = 100 // change

	MaxWalletId := 3 // @@ to get from DB
	awuList := make([]types.AggWltUsg, MaxWalletId+1)

	aggPeriodId, err := db.AggPeriod.InsertAggPeriod(aggStartAt, aggDurationMinutes)
	fmt.Println("InsertAggPeriod() ind: ", aggPeriodId, "  err:", err)

	if err := getWltAggFromDlPkts(aggStartAt, aggDurationMinutes, awuList); err != nil {
		return err // @@ add path
	}

	addPricesWltAgg(awuList, dlPrice) // error does not matter

	addNonPriceFields(awuList, aggStartAt, aggDurationMinutes, aggPeriodId)

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

		// @@ better to be done when the wallet balance is going to be update
		currBalance, err := db.Wallet.GetWalletBalance(int64(k)) // chnge to getBalance  <agg_balance> not the tmp_balance
		if err != nil {
			fmt.Println("GetWalletBalance(): ", currBalance, " || err:", err) //@@
			return err                                                        // @@
		}
		fmt.Println("GetWalletBalance(): wltId:", k, "  curr balance: ", currBalance, " || err:", err) //@@ to remove
		awuList[k].UpdatedBalance = currBalance + awuList[k].BalanceIncrease
		if awuList[k].UpdatedBalance < 0 {
			fmt.Println("Balance under flow") // @@important log
		}
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

func Setup(conf config.MxpConfig) error {

	// initialization from config will be done here

	// calling accounting routine based on time trigger called here

	testAccounting()

	return nil
}
