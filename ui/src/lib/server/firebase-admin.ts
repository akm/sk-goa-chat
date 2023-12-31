import { initializeApp, apps, app } from 'firebase-admin';

// FirebaseAppError: The default Firebase app already exists.
// This means you called initializeApp() more than once without
// providing an app name as the second argument.
// というエラーが発生していたので単純に initializeApp() を呼ぶのではなく
// apps.length で判定し、存在していた場合には app() を呼ぶように修正
const adminApp = apps.length === 0 ? initializeApp() : app();
export const auth = adminApp.auth();
