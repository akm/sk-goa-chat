import { error, type RequestEvent } from '@sveltejs/kit';
import { auth } from '$lib/server/firebase-admin';
import { parse } from 'cookie';
import { sessionKey, clearSession, setSession } from '$lib/server/session';

// See https://kit.svelte.jp/docs/types#public-types-requestevent
//     https://kit.svelte.jp/docs/routing#server-receiving-data
//     https://learn.svelte.jp/tutorial/event

// See 	https://zenn.dev/ssssota/articles/sveltekit-hooks

export const POST = async (event: RequestEvent): Promise<Response> => {
	const { request } = event;
	const reqBody = await request.json();
	if (reqBody.token === undefined) throw error(400, 'token not found');

	const idToken = reqBody.token as string;
	const expiresIn = 5 * 24 * 60 * 60 * 1000; // 5 days in milliseconds
	const sessionCookie = await auth.createSessionCookie(idToken, { expiresIn });

	const resp = new Response();
	setSession(resp, sessionCookie, expiresIn);
	return resp;
};

export const DELETE = async (event: RequestEvent) => {
	const { request } = event;
	const cookies = parse(request.headers.get('cookie') || '');
	try {
		const token = await auth.verifySessionCookie(cookies[sessionKey] ?? '');
		await auth.revokeRefreshTokens(token.sub);
	} catch (err) {
		console.error(err);
		throw err;
	}

	const resp = new Response();
	clearSession(resp);
	return resp;
};
