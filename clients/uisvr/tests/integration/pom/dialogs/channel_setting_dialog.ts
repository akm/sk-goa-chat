import type { Locator, Page } from '@playwright/test';

export class ChannelSettingDialog {
	constructor(readonly page: Page) {}

	get title(): Locator {
		return this.page.locator('h3:text("Channel Settings")');
	}

	inputField(name: 'Name'): Locator {
		return this.page.locator(`label:has-text("${name}") input`);
	}

	button(name: 'Update' | 'Cancel' | 'Delete this channel forever'): Locator {
		const selector = `button:text("${name}")`;
		return this.page.locator(selector);
	}

	closeButton(): Locator {
		return this.page.locator('button[aria-label="Close modal"]');
	}
}
