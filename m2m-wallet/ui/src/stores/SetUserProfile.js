import sessionStore from "./SessionStore";

export default function updateOrganizations(response) {
    /* if(response.body.userProfile.isOrgListModified){
        sessionStore.setOrganizationList(response.body.userProfile.organizations);
    } */
    
    if(true){
        sessionStore.setOrganizationList(response.body.userProfile.organizations);
    } 
    return response;
};