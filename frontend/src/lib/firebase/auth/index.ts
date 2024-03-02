import {
	getAuth,
	connectAuthEmulator
	// setPersistence,
	// browserLocalPersistence
} from 'firebase/auth';
import { app } from '../app';

export const auth = getAuth(app);
// これを行うと、ユーザーがブラウザを閉じてもログイン状態が維持されるはずなんですが、うまく機能してくれませんでした。
// setPersistence(auth, browserLocalPersistence);

// https://firebase.google.com/docs/emulator-suite/connect_auth?hl=ja#web-modular-api
if (import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST) {
	connectAuthEmulator(auth, 'http://' + import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST);
}

export type { MessageKey } from './message_keys';
export { messageKeys, isMessageKey } from './message_keys';
