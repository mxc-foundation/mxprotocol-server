import { EventEmitter } from "events";

import Swagger from "swagger-client";

import sessionStore from "./SessionStore";
import {checkStatus, errorHandler } from "./helpers";
import updateOrganizations from "./SetUserProfile";
import dispatcher from "../dispatcher";


class ProfileStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/profile.swagger.json", sessionStore.getClientOpts());
  }

  getUserOrganizationList(orgId, callbackFunc) {
    this.swagger.then(client => {
      client.apis.InternalService.GetUserOrganizationList({
        orgId
      })
      .then(checkStatus)
      .then(updateOrganizations)
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
        message: "Profile has been " + action,
      },
    });
  }
}

const profileStore = new ProfileStore();
export default profileStore;
