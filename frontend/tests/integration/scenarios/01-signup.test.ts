import { test } from '@playwright/test';
import { signup } from '../steps/signup';
import { signin, signout } from '../steps/signin';
import { saveCookiesTo } from '../steps/auth';
import { foo } from './config';

test('show signin page when go to root', async ({ page, context }) => {
	await signup(page, foo);
	await signout(page);
	await signin(page, foo);

	await saveCookiesTo(context, foo.cookieFile);
});
