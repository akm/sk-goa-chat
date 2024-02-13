import { test } from '@playwright/test';
import { signup } from '../steps/signup';
import { signin, signout } from '../steps/signin';

test('show signin page when go to root', async ({ page }) => {
	const foo = {
		email: 'foo@example.com',
		password: 'Passw0rd!',
		name: 'Foo'
	};

	await signup(page, foo);
	await signout(page);
	await signin(page, foo);
});
