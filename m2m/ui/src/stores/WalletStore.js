import { EventEmitter } from "events";

import Swagger from "swagger-client";

import i18n, { packageNS } from '../i18n';
import sessionStore from "./SessionStore";
import {checkStatus, errorHandler } from "./helpers";
import dispatcher from "../dispatcher";


class WalletStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/wallet.swagger.json", sessionStore.getClientOpts());
  }

  getWalletBalance(orgId, callbackFunc) {
    this.swagger.then(client => {
      client.apis.WalletService.GetWalletBalance({
        orgId,
      })
      .then(checkStatus)
      //.then(updateOrganizations)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  getDlPrice(orgId, callbackFunc) {
    this.swagger.then(client => {
      client.apis.WalletService.GetDlPrice({
        orgId,
      })
      .then(checkStatus)
      //.then(updateOrganizations)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  getWalletUsageHist(orgId, offset, limit, callbackFunc) {
    this.swagger.then(client => {
      client.apis.WalletService.GetWalletUsageHist({
        orgId,
        offset,
        limit
      })
      .then(checkStatus)
      //.then(updateOrganizations)
      .then(resp => {
        callbackFunc(resp.body);
      })
      .catch(errorHandler);
    });
  }

  notify(action) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "success",
        message: `${i18n.t(`${packageNS}:menu.store.balance_has_been`)} ` + action,
      },
    });
  }
}

const walletStore = new WalletStore();
export default walletStore;
