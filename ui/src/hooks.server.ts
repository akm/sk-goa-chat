import type { Handle } from '@sveltejs/kit';
// import { FirebaseError } from 'firebase-admin';

import { auth } from '$lib/server/firebase-admin';
import type { DecodedIdToken } from 'firebase-admin/lib/auth/token-verifier';
import { isFirebaseError } from '$lib/firebase';
import { isMessageKey } from '$lib/firebase/auth';

export const handle: Handle = async ({ event, resolve }) => {
	// get cookies from browser
	const sessionID = event.cookies.get('__session');

	if (!sessionID) {
		// if there is no session load page as normal
		return await resolve(event);
	}

	let decoded: DecodedIdToken;
	try {
		decoded = await auth.verifySessionCookie(sessionID, true);
	} catch (err) {
		if (isFirebaseError(err)) {
			if (isMessageKey(err.code)) {
				switch (err.code) {
					case 'auth/session-cookie-revoked':
					case 'auth/user-not-found':
						return await resolve(event);
				}
			}
		}
		console.error('verifySessionCookie error', err);
		throw err;
	}
	const user = await auth.getUser(decoded.uid);
	event.locals.user = {
		id: user.uid,
		name: user.displayName,
		email: user.email
	};

	// load page as normal
	return await resolve(event);
};
