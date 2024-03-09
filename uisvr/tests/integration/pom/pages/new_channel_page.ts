import type { Locator, Page } from '@playwright/test';

export class NewChannelPage {
	constructor(readonly page: Page) {}

	get title(): Locator {
		return this.page.locator('h3:text("New Channel")');
	}

	inputField(name: 'Name'): Locator {
		return this.page.locator(`label:has-text("${name}") input`);
	}

	button(name: 'Create'): Locator {
		const selector = `button:text("${name}")`;
		return this.page.locator(selector);
	}
}
