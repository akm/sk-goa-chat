import { expect, test } from '@playwright/test';
import type { Page } from '@playwright/test';

import { Sidebar } from '../pom/panes/sidebar';
import { SigninPage } from '../pom/pages/signin_page';

export const signin = async (
	page: Page,
	user: {
		email: string;
		password: string;
	}
) => {
	const signinPage = new SigninPage(page);

	await test.step('signin', async () => {
		const p = signinPage;
		await p.inputField('メールアドレス').fill(user.email);
		await p.inputField('パスワード').fill(user.password);
		await p.button('ログイン').click();
		await expect(signinPage.title).toBeHidden();
	});
};

export const signout = async (page: Page) => {
	const header = new Sidebar(page);

	await test.step('signout', async () => {
		await expect(header.avatarIcon).toBeVisible();
		await header.avatarIcon.click();
		await expect(header.avatarMenu.locator).toBeVisible();
		await header.avatarMenu.signoutButton.click();

		await expect(page).toHaveURL(/signin/);
		await expect(header.title).toBeVisible();
		await expect(header.avatarIcon).toBeHidden();
	});
};
