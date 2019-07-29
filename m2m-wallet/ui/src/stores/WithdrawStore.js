import { EventEmitter } from "events";

import Swagger from "swagger-client";

import sessionStore from "./SessionStore";
import {checkStatus, errorHandler } from "./helpers";
import updateOrganizations from "./SetUserProfile";
import dispatcher from "../dispatcher";


class WithdrawStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/withdraw.swagger.json", sessionStore.getClientOpts());
  }

  getWithdrawFee(money_abbr, orgId, callbackFunc) {
    this.swagger.then(client => {
      client.apis.WithdrawService.GetWithdrawFee({
        money_abbr,
        orgId
      })
      .then(checkStatus)
      //.then(updateOrganizations)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  WithdrawReq(apiWithdrawReqRequest, callbackFunc) {
    this.swagger.then(client => {
      client.apis.WithdrawService.WithdrawReq({
        "money_abbr": apiWithdrawReqRequest.moneyAbbr,
        body: {
          apiWithdrawReqRequest,
        },
      })
      .then(checkStatus)
      //.then(updateOrganizations)
      .then(resp => {
        this.notify("updated");
        this.emit("withdraw");
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
        message: "Withdrawal succeeded"
      },
    });
  }
}

const withdrawStore = new WithdrawStore();
export default withdrawStore;
