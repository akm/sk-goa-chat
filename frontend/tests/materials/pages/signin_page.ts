import type { Locator, Page } from '@playwright/test';

export class SigninPage {
	constructor(readonly page: Page) {}

	get title(): Locator {
		return this.page.locator('h3:text("ログイン")');
	}

	inputField(name: 'メールアドレス' | 'パスワード'): Locator {
		return this.page.locator(`label:has-text("${name}") input`);
	}

	button(name: 'ログイン' | 'アカウント登録'): Locator {
		const selector =
			name === 'ログイン' ? 'button:text("ログイン")' : 'a[role="button"]:text("アカウント登録")';
		return this.page.locator(selector);
	}
}
