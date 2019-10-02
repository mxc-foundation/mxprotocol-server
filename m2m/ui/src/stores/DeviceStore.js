import { EventEmitter } from "events";

import Swagger from "swagger-client";

import sessionStore from "./SessionStore";
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
            for(var i=0;i<10;i++){

            resp.body.devProfile.push({rownum:i, name: '0x1234567',
            lastSeenAt: '1x1234',
            mode: 'Ether',
            devEui: '0.0000000000123456',
            createdAt: '2019-09-01 112345674345'});
            }
            resp.body.count = 10
            console.log(resp.body);
 
            callbackFunc(resp.body);
          })
          .catch(errorHandler);
        });
    }

    /* getDeviceHistory(orgId, offset, limit, callbackFunc) {    
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
    } */

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
            message: "Device(s) has been " + action,
        },
    });
  }
}

const deviceStore = new DeviceStore();
export default deviceStore;
