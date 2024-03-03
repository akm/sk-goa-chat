import { test } from '@playwright/test';
import { signup } from '../steps/signup';
import { signin, signout } from '../steps/signin';
import { foo } from './config';

test('show signin page when go to root', async ({ page }) => {
	await signup(page, foo);
	await signout(page);
	await signin(page, foo);
});
