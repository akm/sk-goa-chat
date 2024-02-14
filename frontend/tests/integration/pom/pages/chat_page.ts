import type { Locator, Page } from '@playwright/test';

export class ChatPage {
	constructor(readonly page: Page) {}

	title(name: string): Locator {
		return this.page.locator(`h3:text("${name}")`);
	}

	get configButton(): Locator {
		return this.page.locator('button:has(svg[aria-label="cog outline"])');
	}

	get textarea(): Locator {
		return this.page.locator('textarea');
	}

	button(name: 'Read previous messages' | 'Read next messages' | 'Send'): Locator {
		const selector = `button:text("${name}")`;
		return this.page.locator(selector);
	}

	get messagePane(): Locator {
		return this.page.locator('[data-testid="message_pane"]');
	}
}
