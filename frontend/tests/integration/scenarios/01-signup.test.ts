import { test } from '@playwright/test';
import { signup } from '../steps/signup';
import { signin, signout } from '../steps/signin';
import { saveIndexedDBDataTo } from '../steps/indexeddb';
import { foo, firebaseIndexedDBConfig } from './config';

test('show signin page when go to root', async ({ page }) => {
	await signup(page, foo);
	await signout(page);
	await signin(page, foo);

	const { dbName, objectStoreName } = firebaseIndexedDBConfig;
	await saveIndexedDBDataTo(page, dbName, objectStoreName, foo.credentialFilePath);
	await page.pause();
});
