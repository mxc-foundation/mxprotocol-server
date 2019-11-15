import { EventEmitter } from "events";

import Swagger from "swagger-client";
import { checkStatus, errorHandler, errorHandlerLogin } from "./helpers";
//import updateOrganizations from "./SetUserProfile";

class SessionStore extends EventEmitter {
  constructor() {
    super();
    this.client = null;
    this.user = null;
    this.organizations = [];
    this.settings = {};
    this.branding = {};

    this.swagger = Swagger("/swagger/profile.swagger.json", this.getClientOpts());

    this.swagger.then(client => {
      this.client = client;

      /* if (this.getToken() !== null) {
        this.fetchProfile(() => {});
      } */
    });
  }

  getClientOpts() {
    return {
      requestInterceptor: (req) => {
        if (this.getToken() !== null) {
          req.headers["Grpc-Metadata-Authorization"] = "Bearer " + this.getToken();
        }
        return req;
      },
    }
  }

  setToken(token) {
    localStorage.setItem("jwt", token);
  }

  getToken() {
    return localStorage.getItem("jwt");
  }

  setSupportedLanguages(languages) {
    localStorage.setItem("languages-supported", JSON.stringify(languages));
  }

  getSupportedLanguages() {
    return JSON.parse(localStorage.getItem("languages-supported"));
  }

  setLanguage(language) {
    localStorage.setItem("language", JSON.stringify(language));
  }

  getLanguage() {
    return JSON.parse(localStorage.getItem("language"));
  }

  setUsername(username) {
    localStorage.setItem("username", username);
  }

  getUsername() {
    return localStorage.getItem("username");
  }

  getOrganizationID() {
    const orgID = localStorage.getItem("organizationID");
    if (orgID === "") {
      return null;
    }

    return orgID;
  }

  setOrganizationID(id) {
    localStorage.setItem("organizationID", id);
    this.emit("organization.change");
  }

  getOrganizationName() {
    const orgName = localStorage.getItem("organizationName");
    if (orgName === "") {
      return null;
    }

    return orgName;
  }

  setOrganizationName(name) {
    localStorage.setItem("organizationName", name);
  }

  setOrganizationList(organizations) {
    let organizationList = null;
    
    if(organizations.length > 0){
      organizationList = organizations.map((o, i) => { 
      return {label: o.organizationName, value: o.organizationID}});
    }
    
    localStorage.setItem("organizationList", JSON.stringify(organizationList));
    this.emit("organizationList.change");
    //this.emit("organizationList.change");
  }

  getOrganizationList() {
    //debugger
    const organizationList = localStorage.getItem("organizationList");
    if (!organizationList) {
      return [];
    }

    return JSON.parse(organizationList);
  }

  getUserOrganizationList(orgId, callbackFunc) {
    this.swagger.then(client => {
      client.apis.InternalService.GetUserOrganizationList({
        orgId
      })
      .then(checkStatus)
      //.then(updateOrganizations)
      .then(resp => {
        const organizations = resp.body.organizations;
        let organizationList = null;
    
        if(organizations.length > 0){
          organizationList = organizations.map((o, i) => { 
          return {label: o.organizationName, value: o.organizationID}});
        }else{
          organizationList = {};
        }
    
    
        callbackFunc(organizationList);
      })
      .catch(errorHandler);
    });
  }

  getUser() {
    return JSON.parse(localStorage.getItem('user'));
  }

  getSettings() {
    return this.settings;
  }

  isAdmin() {
    this.user = this.getUser()
    if (this.user === undefined || this.user === null) {
      return false;
    }
    return this.user.isAdmin;
  }

  isOrganizationAdmin(organizationID) {
    for (let i = 0; i < this.organizations.length; i++) {
      if (this.organizations[i].organizationID === organizationID) {
        return this.organizations[i].isAdmin;
      }
    }
  }

  initProfile(data) {

    const { jwt, orgId, orgName, username, loraHostUrl } = data;
    
    if(jwt === "" || orgId === "" || orgId === undefined){
      window.location.replace(loraHostUrl);
    }
    
    this.setToken(jwt);
    this.setUsername(username);
    this.setOrganizationID(orgId);
    this.setOrganizationName(orgName);
    this.fetchProfile()
  }

  login(login, callBackFunc) {
    this.swagger.then(client => {
      client.apis.InternalService.Login({body: login})
        .then(checkStatus)
        .then(resp => {
          if(resp.body.jwt === ""){
            callBackFunc(false);  
          }else{
            callBackFunc(true);
          }
          //this.fetchProfile(callBackFunc);
        })
        .catch(errorHandlerLogin);
    });
  }

  logout(callBackFunc) {
    localStorage.clear();
    this.user = null;
    this.organizations = [];
    this.settings = {};
    this.emit("change");
    callBackFunc();
  }

  fetchProfile(callBackFunc) {
    const orgId = this.getOrganizationID()
    this.swagger.then(client => {
      client.apis.InternalService.GetUserProfile({
        orgId: orgId
      })
        .then(checkStatus)
        .then(resp => {
          localStorage.setItem("user", JSON.stringify(resp.obj.userProfile.user))

          if(resp.obj.organizations !== undefined) {
            this.organizations = resp.obj.organizations;
          }

          if(resp.obj.settings !== undefined) {
            this.settings = resp.obj.settings;
          }

          this.emit("change");
          //callBackFunc();
        })
        .catch(errorHandler);
    });
  }

  globalSearch(search, limit, offset, callbackFunc) {
    this.swagger.then(client => {
      client.apis.InternalService.GlobalSearch({
        search: search,
        limit: limit,
        offset: offset,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
      });
  }

  getBranding(callbackFunc) {
    return false;
    /* this.swagger.then(client => {
      client.apis.InternalService.Branding({})
        .then(checkStatus)
        .then(resp => {
          callbackFunc(resp.obj);
        })
        .catch(errorHandler);
    }); */
  }
}

const sessionStore = new SessionStore();
export default sessionStore;
