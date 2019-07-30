
import SessionStore from '../stores/SessionStore';

export const LoraUrl = SessionStore.getLoraHostUrl();

export function redirectToLora() {
    window.location.replace(LoraUrl);
}