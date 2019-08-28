import { EventEmitter } from "events";

import Swagger from "swagger-client";

import sessionStore from "./SessionStore";
import {checkStatus, errorHandler } from "./helpers";
import dispatcher from "../dispatcher";


class SupernodeStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/super_node.swagger.json", sessionStore.getClientOpts());
  }

  getSuperNodeActiveMoneyAccount(money_abbr, callbackFunc) {
    this.swagger.then(client => {
      client.apis.SuperNodeService.GetSuperNodeActiveMoneyAccount({
        money_abbr
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.body);
      })
      .catch(errorHandler);
    });
  }

  addSuperNodeMoneyAccount(req, callbackFunc) {
    this.swagger.then(client => {
      client.apis.SuperNodeService.AddSuperNodeMoneyAccount({
        "money_abbr": req.moneyAbbr,
        body: {
            moneyAbbr: req.moneyAbbr,
            accountAddr: req.createAccount
        },
      })
      .then(checkStatus)
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
        message: "user has been " + action,
      },
    });
  }
}

const supernodeStore = new SupernodeStore();
export default supernodeStore;