import type { Handle } from '@sveltejs/kit';
// import { FirebaseError } from 'firebase-admin';

import { auth } from '$lib/server/firebase-admin';

export const handle: Handle = async ({ event, resolve }) => {
	const uid = event.request.headers.get('X-UID');
	if (!uid) {
		return await resolve(event);
	}

	const user = await auth.getUser(uid);

	event.locals.user = {
		id: user.uid,
		name: user.displayName,
		email: user.email
	};

	// load page as normal
	return await resolve(event);
};
