import { getAuth, connectAuthEmulator } from 'firebase/auth';
import { app } from '../app';

export const auth = getAuth(app);

// https://firebase.google.com/docs/emulator-suite/connect_auth?hl=ja#web-modular-api
if (import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST) {
	connectAuthEmulator(auth, 'http://' + import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST);
}

export type { MessageKey } from './message_keys';
export { messageKeys, isMessageKey } from './message_keys';
