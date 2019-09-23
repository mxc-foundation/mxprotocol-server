
//import SessionStore from '../stores/SessionStore';

//export const LoraUrl = SessionStore.getLoraHostUrl();
//orgId is 0 which means current user is super_admin
export const SUPER_ADMIN = '0';

export function redirectToLora() {
    let host = process.env.REACT_APP_LORA_APP_SERVER;
    const origin = window.location.origin;
    
    if(origin.includes(process.env.REACT_APP_SUBDOM_M2M)){
        host = origin.replace(process.env.REACT_APP_SUBDOM_M2M, process.env.REACT_APP_SUBDOM_LORA);
    }
    
    return host
}