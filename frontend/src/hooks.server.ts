import type { Handle } from '@sveltejs/kit';
// import { FirebaseError } from 'firebase-admin';

import { auth } from '$lib/server/firebase-admin';
import type { DecodedIdToken } from 'firebase-admin/auth';
import { isFirebaseError } from '$lib/firebase';
import { isMessageKey } from '$lib/firebase/auth';

export const handle: Handle = async ({ event, resolve }) => {
	const idToken = event.request.headers.get('X-ID-TOKEN');
	if (!idToken) {
		return await resolve(event);
	}
	event.locals.idToken = idToken;

	let decodedToken: DecodedIdToken;
	try {
		decodedToken = await auth.verifyIdToken(idToken);
		// decoded の中身の例
		// {
		// 	email: 'foo1@example.com',
		// 	email_verified: false,
		// 	auth_time: 1704273607,
		// 	user_id: 'nf5LAN4rZdK3vgNQPikElX1I5Dri',
		// 	firebase: { identities: { email: [Array] }, sign_in_provider: 'password' },
		// 	iat: 1704273607,
		// 	exp: 1704705607,
		// 	aud: 'sk-goa-chat',
		// 	iss: 'https://session.firebase.google.com/sk-goa-chat',
		// 	sub: 'nf5LAN4rZdK3vgNQPikElX1I5Dri',
		// 	uid: 'nf5LAN4rZdK3vgNQPikElX1I5Dri'
		// }
		//
		// firebase.identities.email は [ 'foo1@example.com' ]
	} catch (err) {
		if (isFirebaseError(err)) {
			if (isMessageKey(err.code)) {
				switch (err.code) {
					case 'auth/session-cookie-revoked':
					case 'auth/user-not-found':
						event.locals.user = undefined;
						event.locals.idToken = undefined;
						return await resolve(event);
				}
			}
		}
		console.error('verifySessionCookie error', err);
		throw err;
	}

	const user = await auth.getUser(decodedToken.uid);
	event.locals.user = {
		id: user.uid,
		name: user.displayName,
		email: user.email
	};

	// load page as normal
	return await resolve(event);
};
