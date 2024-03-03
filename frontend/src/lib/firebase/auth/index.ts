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

export let authReady = false;

auth.onAuthStateChanged((user) => {
	if (!user) {
		authReady = false;
		console.log('src/lib/firebase/auth/index.ts/auth.onAuthStateChanged no user');
		return;
	}
	authReady = true;
	console.log('src/lib/firebase/auth/index.ts/auth.onAuthStateChanged user: ', user);
});

const waitAuth =
	(expected: boolean) =>
	async (opts?: { attempts?: number; interval?: number }): Promise<void> => {
		const attempts = opts?.attempts ?? 30;
		const interval = opts?.interval ?? 100;
		for (let i = 0; i < attempts; i++) {
			if (authReady === expected) {
				const name = expected ? 'waitUntilSignedIn' : 'waitUntilSignedOut';
				console.log(`${name} success`);
				return;
			}
			await new Promise((resolve) => setTimeout(resolve, interval));
		}
		throw new Error(`authReady timeout expected: ${expected}`);
	};

export const waitUntilSignedIn = waitAuth(true);
export const waitUntilSignedOut = waitAuth(false);

// https://firebase.google.com/docs/emulator-suite/connect_auth?hl=ja#web-modular-api
if (import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST) {
	connectAuthEmulator(auth, 'http://' + import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST);
}

export type { MessageKey } from './message_keys';
export { messageKeys, isMessageKey } from './message_keys';
