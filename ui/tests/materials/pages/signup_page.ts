import type { Locator, Page } from '@playwright/test';

export class SignupPage {
	constructor(readonly page: Page) {}

	get title(): Locator {
		return this.page.locator('h3:text("アカウント登録")');
	}

	inputField(name: 'メールアドレス' | 'パスワード' | 'アカウント名'): Locator {
		return this.page.locator(`label:has-text("${name}") input`);
	}

	button(name: 'アカウント登録' | 'ログイン'): Locator {
		const selector =
			name === 'アカウント登録'
				? 'button:text("アカウント登録")'
				: 'a[role="button"]:text("ログイン")';
		return this.page.locator(selector);
	}
}
