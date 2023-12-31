// See https://firebase.google.com/docs/admin/setup?hl=ja
import { initializeApp, getApp, getApps } from 'firebase-admin/app';
import { getAuth } from 'firebase-admin/auth';

// FirebaseAppError: The default Firebase app already exists.
// This means you called initializeApp() more than once without
// providing an app name as the second argument.
// というエラーが npm run dev での実行時に発生していたので単純に
// その場合は initializeApp() を呼ぶのではなく、apps.length で判定し、
// 存在していた場合には app() を呼ぶように修正
// See https://vitejs.dev/guide/env-and-mode
const adminApp = (() => {
	if (import.meta.env.DEV) {
		return getApps().length === 0 ? initializeApp() : getApp();
	} else {
		return initializeApp();
	}
})();

export const auth = getAuth(adminApp);
