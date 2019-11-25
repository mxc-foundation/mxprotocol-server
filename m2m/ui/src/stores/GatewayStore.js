import { EventEmitter } from "events";

import Swagger from "swagger-client";

import i18n, { packageNS } from '../i18n';
import sessionStore from "./SessionStore";
import {checkStatus, errorHandler } from "./helpers";
import dispatcher from "../dispatcher";


class GatewayStore extends EventEmitter {
    constructor() {
        super();
        this.swagger = new Swagger("/swagger/gateway.swagger.json", sessionStore.getClientOpts());
    }

    getGatewayList(orgId, offset, limit, callbackFunc) {
        this.swagger.then(client => {
          client.apis.GatewayService.GetGatewayList({
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
    
      getGatewayProfile(gwId, callbackFunc) {
        this.swagger.then(client => {
            client.apis.GatewayService.GetGatewayProfile({
            gwId,
            })
            .then(checkStatus)
            .then(resp => {
            callbackFunc(resp.obj);
            })
            .catch(errorHandler);
        });
    }

    getGatewayHistory(orgId, gwId, offset, limit, callbackFunc) {    
        this.swagger.then(client => {
            client.apis.GatewayService.GetGatewayHistory({
            orgId,
            gwId,
            offset,
            limit
            })
            .then(checkStatus)
            .then(resp => {
                callbackFunc(resp.body);
            })
            .catch(errorHandler);
        });
    }

    setGatewayMode(orgId, gwId, gwMode, callbackFunc) {
        this.swagger.then(client => {
        client.apis.GatewayService.SetGatewayMode({
            "orgId": orgId,
            "gwId": gwId,
            body: {
                orgId,
                gwId,
                gwMode
            },
        })
        .then(checkStatus)
        .then(resp => {
            this.emit("update");
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
            message: `${i18n.t(`${packageNS}:menu.store.gateway_has_been`)} ` + action,
        },
    });
  }
}

const gatewayStore = new GatewayStore();
export default gatewayStore;
