import { test } from '@playwright/test';
import { signup } from '../steps/signup';
import { signin, signout } from '../steps/signin';
import { getFoo } from './config';

test('show signin page when go to root', async ({ page }, workers) => {
	const foo = getFoo(workers.workerIndex);
	await signup(page, foo);
	await signout(page);
	await signin(page, foo);
});
