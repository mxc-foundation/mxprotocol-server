import { EventEmitter } from "events";

import Swagger from "swagger-client";

import sessionStore from "./SessionStore";
import i18n, { packageNS } from '../i18n';
import {checkStatus, errorHandler } from "./helpers";
import dispatcher from "../dispatcher";


class DeviceStore extends EventEmitter {
    constructor() {
        super();
        this.swagger = new Swagger("/swagger/device.swagger.json", sessionStore.getClientOpts());
    }

    getDeviceList(orgId, offset, limit, callbackFunc) {
        this.swagger.then(client => {
          client.apis.DeviceService.GetDeviceList({
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

    getDeviceHistory(orgId, offset, limit, callbackFunc) {    
        this.swagger.then(client => {
            client.apis.DeviceService.GetDeviceHistory({
            orgId,
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

    setDeviceMode(orgId, devId, devMode, callbackFunc) {
        this.swagger.then(client => {
        client.apis.DeviceService.SetDeviceMode({
            "orgId": orgId,
            "devId": devId,
            body: {
                orgId,
                devId,
                devMode
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
            // TODO - convert to i18n
            message: `${i18n.t(`${packageNS}:menu.store.devices_has_been`)} ` + action,
        },
    });
  }
}

const deviceStore = new DeviceStore();
export default deviceStore;
