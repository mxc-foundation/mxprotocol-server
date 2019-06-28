import { EventEmitter } from "events";

import Swagger from "swagger-client";

import sessionStore from "./SessionStore";
import dispatcher from "../dispatcher";


class HistoryStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/withdraw.swagger.json", sessionStore.getClientOpts());
  }

  list(search, limit, offset, callbackFunc) {
    /* this.swagger.then((client) => {
      client.apis.UserService.List({
        search: search,
        limit: limit,
        offset: offset,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    }); */
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
