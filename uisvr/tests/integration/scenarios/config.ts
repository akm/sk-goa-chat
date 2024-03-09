export const getFoo = (index: number) => {
	return {
		email: `foo${index}@example.com`,
		password: 'Passw0rd!',
		name: `Foo${index}`,
		cookieFile: `tests/integration/tmp/cookies-foo${index}.json`
	};
};

export const getBar = (index: number) => {
	return {
		email: `bar${index}@example.com`,
		password: 'Passw0rd!',
		name: `Bar${index}`,
		cookieFile: `tests/integration/tmp/cookies-bar${index}.json`
	};
};
