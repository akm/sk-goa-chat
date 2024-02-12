import { expect, test } from '@playwright/test';
import { Sidebar } from './pom/panes/sidebar';
import { SigninPage } from './pom/pages/signin_page';
import { SignupPage } from './pom/pages/signup_page';

test('show signin page when go to root', async ({ page }) => {
	const header = new Sidebar(page);
	const signinPage = new SigninPage(page);
	const signupPage = new SignupPage(page);

	await test.step('open top page before signin', async () => {
		await page.goto('/');
		await expect(page).toHaveURL(/signin/);
		await expect(header.title).toBeVisible();
		await expect(signinPage.title).toBeVisible();
		await expect(header.avatarIcon).toBeHidden();
	});

	await test.step('signup', async () => {
		await signinPage.button('アカウント登録').click();
		await expect(header.avatarIcon).toBeHidden();
		await expect(signupPage.title).toBeVisible();
		const p = signupPage;
		await p.inputField('メールアドレス').fill('foo@example.com');
		await p.inputField('パスワード').fill('Passw0rd!');
		await p.inputField('アカウント名').fill('Foo');
		await p.button('アカウント登録').click();
		await expect(signupPage.title).toBeHidden();
		await expect(signinPage.title).toBeHidden();
	});

	await test.step('signout', async () => {
		await expect(header.avatarIcon).toBeVisible();
		await header.avatarIcon.click();
		await expect(header.avatarMenu.locator).toBeVisible();
		await header.avatarMenu.signoutButton.click();

		await expect(page).toHaveURL(/signin/);
		await expect(header.title).toBeVisible();
		await expect(signinPage.title).toBeVisible();
		await expect(header.avatarIcon).toBeHidden();
	});

	await test.step('signin', async () => {
		const p = signinPage;
		await p.inputField('メールアドレス').fill('foo@example.com');
		await p.inputField('パスワード').fill('Passw0rd!');
		await p.button('ログイン').click();
		await expect(signupPage.title).toBeHidden();
		await expect(signinPage.title).toBeHidden();
	});
});
