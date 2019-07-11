import { EventEmitter } from "events";

import Swagger from "swagger-client";

import sessionStore from "./SessionStore";
import {checkStatus, errorHandler } from "./helpers";
import dispatcher from "../dispatcher";


class HistoryStore extends EventEmitter {
  constructor() {
    super();
    this.topupSwagger = new Swagger("/swagger/topup.swagger.json", sessionStore.getClientOpts());
    this.withdrawSwagger = new Swagger("/swagger/withdraw.swagger.json", sessionStore.getClientOpts());
    this.walletSwagger = new Swagger("/swagger/wallet.swagger.json", sessionStore.getClientOpts());
    this.moneySwagger = new Swagger("/swagger/ext_account.swagger.json", sessionStore.getClientOpts());
  }

  getTopUpHistory(orgId, limit, offset, callbackFunc) {
    this.topupSwagger.then((client) => {      
      client.apis.TopupService.List({
        orgId,
        limit,
        offset,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }
  
  getWithdrawHistory(money_abbr, orgId, limit, offset, callbackFunc) {
    this.withdrawSwagger.then((client) => {      
      client.apis.WithdrawService.GetWithdrawHistory({
        money_abbr,
        orgId,
        limit,
        offset,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  getVmxcTxHistory(orgId, limit, offset, callbackFunc) {
    this.walletSwagger.then((client) => {      
      client.apis.WalletService.GetVmxcTxHistory({
        orgId,
        limit,
        offset,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  getChangeMoneyAccountHistory(money_abbr, orgId, limit, offset, callbackFunc) {
    this.moneySwagger.then((client) => {      
      client.apis.MoneyService.GetChangeMoneyAccountHistory({
        money_abbr,
        orgId,
        limit,
        offset,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  notify(action) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "success",
        message: "user has been " + action,
      },
    });
  }
}

const historyStore = new HistoryStore();
export default historyStore;
