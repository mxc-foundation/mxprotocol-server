package accounting

import (
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func performAccounting(aggDurationMinutes int64, dlPrice float64) error {

	log.WithFields(log.Fields{
		"dl_price": dlPrice,
	}).Info("Accounting routine started!")

	// aggPeriodId, err := db.AggPeriod.InsertAggPeriod(aggStartAt, aggDurationMinutes, execTime)

	latestAccountedDlPktId := db.AggPeriod.GetLatestAccountedDlPktId()
	aggPeriodId, err := db.AggPeriod.InsertAggPeriod(latestAccountedDlPktId, aggDurationMinutes)

	if err != nil {
		return errors.Wrap(err, "accounting/performAccounting: Unable to start accounting")
	}
	log.Info("accounting/ Aggregation Period: ", aggPeriodId)

	MaxWalletId, errMaxWalletId := db.Wallet.GetMaxWalletId()
	if errMaxWalletId != nil {
		return errors.Wrap(errMaxWalletId, "accounting/performAccounting Unable to start accounting")
	}

	awuList := make([]types.AggWltUsg, MaxWalletId+1)
	if err := getWltAggFromDlPkts(aggStartAt, aggDurationMinutes, awuList); err != nil {
		return err
	}

	addPricesWltAgg(awuList, dlPrice)

	addNonPriceFields(awuList, aggStartAt, aggDurationMinutes, aggPeriodId)

	walletIdSuperNode, errWltId := db.Wallet.GetWalletIdSuperNode()
	if errWltId != nil {
		return errors.Wrap(errWltId, "accounting/performAccounting  Unable to write accounting in DB; unable to get superNodeAccount")
	}

	log.WithFields(log.Fields{
		"Accounting  aggPeriodId": aggPeriodId,
	}).Info("Accounting calculations is done; writing in the DB started!")

	if err := putInDbAggWltUsg(awuList, walletIdSuperNode); err != nil {
		return errors.Wrap(errWltId, "accounting/performAccounting")
	}

	log.WithFields(log.Fields{
		"Accounting  aggPeriodId": aggPeriodId,
	}).Info("Accounting performed successfully!")

	return nil

}

func getWltAggFromDlPkts(aggStartAt time.Time, aggDurationMinutes int64, awuList []types.AggWltUsg) error {

	if wltIds, cnts, err := db.DlPacket.GetAggDlPktDeviceWallet(aggStartAt, aggDurationMinutes); true {
		// fmt.Println("GetAggDlPktDeviceWallet   wltIds: ", wltIds, "  cnts: ", cnts, "   err: ", err)
		if err != nil {
			return errors.Wrap(err, "accounting/getWltAggFromDlPkts")
		}
		if len(wltIds) != len(cnts) {
			return errors.Wrap(err, "accounting/getWltAggFromDlPkts  Inequal length of arrays wltIds, Cnts GetAggDlPktDeviceWallet")
		}
		for k, v := range wltIds {
			awuList[v].DlCntDv = cnts[k]
		}
	}

	if wltIds, cnts, err := db.DlPacket.GetAggDlPktGatewayWallet(aggStartAt, aggDurationMinutes); true {
		// fmt.Println("GetAggDlPktGatewayWallet   wltIds: ", wltIds, "  cnts: ", cnts, "   err: ", err)
		if err != nil {
			return errors.Wrap(err, "accounting/getWltAggFromDlPkts")
		}
		if len(wltIds) != len(cnts) {
			return errors.Wrap(err, "accounting/getWltAggFromDlPkts  Inequal length of arrays wltIds, Cnts GetAggDlPktGatewayWallet")
		}
		for k, v := range wltIds {
			awuList[v].DlCntGw = cnts[k]
		}
	}

	if wltIds, cnts, err := db.DlPacket.GetAggDlPktFreeWallet(aggStartAt, aggDurationMinutes); true {
		// fmt.Println("GetAggDlPktFreeWallet   wltIds: ", wltIds, "  cnts: ", cnts, "   err: ", err)
		if err != nil {
			return errors.Wrap(err, "accounting/getWltAggFromDlPkts")
		}
		if len(wltIds) != len(cnts) {
			return errors.Wrap(err, "accounting/GetAggDlPktFreeWallet  Inequal length of arrays wltIds, Cnts GetAggDlPktGatewayWallet")

		}
		for k, v := range wltIds {
			awuList[v].DlCntDvFree = cnts[k]
			awuList[v].DlCntGwFree = cnts[k]
		}
	}
	return nil
}

func addPricesWltAgg(awuList []types.AggWltUsg, dlPrice float64) {
	for k, v := range awuList {

		if v == (types.AggWltUsg{}) {
			continue
		}

		awuList[k].Spend = float64(v.DlCntDv-v.DlCntDvFree) * dlPrice
		awuList[k].Income = float64(v.DlCntGw-v.DlCntGwFree) * dlPrice
		awuList[k].BalanceIncrease = awuList[k].Income - awuList[k].Spend

	}
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

	for _, v := range awuList {

		if v == (types.AggWltUsg{}) {
			continue
		}

		insertedAggWltUsgId, errIns := db.AggWalletUsage.InsertAggWltUsg(v)
		if errIns != nil {
			log.WithFields(log.Fields{
				"AggWltUsg": v,
			}).WithError(errIns).Error("accounting/putInDbAggWltUsg impossible to write in DB InsertAggWltUsg ")
		}

		_, err := db.AggWalletUsage.ExecAggWltUsgPayments(types.InternalTx{
			FkWalletSender: walletIdSuperNode,
			FkWalletRcvr:   v.FkWallet,
			PaymentCat:     string(types.DOWNLINK_AGGREGATION),
			TxInternalRef:  insertedAggWltUsgId,
			Value:          v.BalanceIncrease,
			TimeTx:         time.Now().UTC(),
		})
		if err != nil {
			log.WithFields(log.Fields{
				"AggWltUsg": v,
			}).WithError(errIns).Error("accounting/putInDbAggWltUsg impossible to write in DB ExecAggWltUsgPayments ")
		}

		syncTmpBalance(v.FkWallet)

	}
	return nil
}
