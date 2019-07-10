import { EventEmitter } from "events";

import Swagger from "swagger-client";

import sessionStore from "./SessionStore";
import {checkStatus, errorHandler } from "./helpers";
import dispatcher from "../dispatcher";


class MoneyStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/ext_account.swagger.json", sessionStore.getClientOpts());
  }

  getActiveMoneyAccount(money_abbr, org_id, callbackFunc) {
    this.swagger.then(client => {
      client.apis.MoneyService.GetActiveMoneyAccount({
        money_abbr,
        org_id,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  modifyMoneyAccount(req, callbackFunc) {
    this.swagger.then(client => {
      client.apis.MoneyService.ModifyMoneyAccount({
        "money_abbr": req.money_abbr,
        body: {
            apiModifyMoneyAccountRequest: req,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.notify("updated");
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

const moneyStore = new MoneyStore();
export default moneyStore;
