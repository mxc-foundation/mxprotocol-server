import { EventEmitter } from "events";

import Swagger from "swagger-client";

import sessionStore from "./SessionStore";
import {checkStatus, errorHandler } from "./helpers";
import dispatcher from "../dispatcher";


class ServerInfoStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/server.swagger.json", sessionStore.getClientOpts());
  }

  getVersion(callbackFunc) {
    this.swagger.then(client => {
      client.apis.ServerInfoService.GetVersion()
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.data);
      })
      .catch(errorHandler);
    });
  }

  notify(action) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "success",
        message: "server has been " + action,
      },
    });
  }
}

const profileStore = new ServerInfoStore();
export default profileStore;
