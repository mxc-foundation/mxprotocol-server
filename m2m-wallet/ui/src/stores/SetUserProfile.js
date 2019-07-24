import sessionStore from "./SessionStore";
import hash from "object-hash";

export default function updateOrganizations(response) {
    
    const organizationList = response.body.userProfile.organizations;
    
    if(!organizationList){
        return false;
    }
    
    if(sessionStore.getOrganizationList() !== null){
        console.log('updateOrganizations', hash(organizationList));
        console.log('updateOrganizations', hash(response.body.userProfile.organizations));
        if(hash(sessionStore.getOrganizationList()) !== hash(organizationList)){
            sessionStore.setOrganizationList(organizationList);
        }
    }else{
        sessionStore.setOrganizationList(organizationList);
    }
     
    return response;
};