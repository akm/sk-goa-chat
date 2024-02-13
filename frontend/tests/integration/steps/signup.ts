import { expect, test } from '@playwright/test';
import type { Page } from '@playwright/test';
import { Sidebar } from '../pom/panes/sidebar';
import { SigninPage } from '../pom/pages/signin_page';
import { SignupPage } from '../pom/pages/signup_page';

export const signup = async (
	page: Page,
	user: {
		email: string;
		password: string;
		name: string;
	}
) => {
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
		await p.inputField('メールアドレス').fill(user.email);
		await p.inputField('パスワード').fill(user.password);
		await p.inputField('アカウント名').fill(user.name);
		await p.button('アカウント登録').click();
		await expect(signupPage.title).toBeHidden();
		await expect(signinPage.title).toBeHidden();
	});
};
