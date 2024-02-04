import { serialize } from 'cookie';

export const sessionKey = 'session_id';

export const setSession = (resp: Response, sessionCookie: string, expiresInMS: number) => {
	resp.headers.set(
		'set-cookie',
		serialize('session_id', sessionCookie, {
			httpOnly: true,
			path: '/',
			secure: true,
			sameSite: 'lax',
			maxAge: expiresInMS / 1000
		})
	);
};

export const clearSession = (resp: Response) => {
	resp.headers.set(
		'set-cookie',
		serialize('session_id', 'revoked', {
			maxAge: 0,
			path: '/'
		})
	);
};
