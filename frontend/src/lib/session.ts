import { POST, DELETE } from './openapi_client';

export const createSession = async (idToken: string) => {
	const { error, response } = await POST('/api/session', {
		params: { header: { 'Content-Type': 'application/json' } },
		body: { id_token: idToken }
	});
	console.log('response', response);
	if (error) {
		console.log('createSession error', error);
		throw error;
	}
};

export const deleteSession = async () => {
	const { error, response } = await DELETE('/api/session', {});
	console.log('response', response);
	if (error) {
		console.log('deleteSession error', error);
		throw error;
	}
};
