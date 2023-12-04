import { serialize } from 'cookie';

export const sessionKey = '__session';

export const setSession = (resp: Response, sessionCookie: string, expiresInMS: number) => {
	resp.headers.set(
		'set-cookie',
		serialize('__session', sessionCookie, {
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
		serialize('__session', 'revoked', {
			maxAge: 0,
			path: '/'
		})
	);
};
