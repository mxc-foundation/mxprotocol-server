import sessionStore from "./SessionStore";
import hash from "object-hash";

export default function updateOrganizations(response) {
    const organizationList = response.body.organizations;
    console.log(organizationList);
    if(!organizationList){
        return false;
    }
    
    if(sessionStore.getOrganizationList() !== null){
        if(hash(sessionStore.getOrganizationList()) !== hash(organizationList)){
            sessionStore.setOrganizationList(organizationList);
        }
    }else{
        sessionStore.setOrganizationList(organizationList);
    }
     
    return response; 
};