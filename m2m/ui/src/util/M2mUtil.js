
import SessionStore from '../stores/SessionStore';

export const LoraUrl = SessionStore.getLoraHostUrl();
//org_id is 0 which means current user is super_admin
export const SUPER_ADMIN = '0';

export function redirectToLora() {
    window.location.replace(process.env.REACT_APP_LORA_APP_SERVER);
}