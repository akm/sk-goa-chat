export const createSession = async (idToken: string) => {
	const response = await fetch('/api/session', {
		method: 'POST',
		body: JSON.stringify({ token: idToken }),
		headers: { 'Content-Type': 'application/json' }
	});
	console.log('response', response);
	if (response.status !== 200) {
		throw new Error('Failed to login');
	}
};

export const deleteSession = async () => {
	const response = await fetch('/api/session', { method: 'DELETE' });
	if (response.status !== 200) {
		throw new Error('Failed to logout');
	}
};
