import type { Locator, Page } from '@playwright/test';

export class Sidebar {
	readonly locator: Locator;
	constructor(page: Page) {
		this.locator = page.locator('nav[data-testid="header_nav"]');
	}

	get title(): Locator {
		return this.locator.locator('a span:text("SK Goa Chat")');
	}

	get avatarIcon(): Locator {
		return this.locator.locator('#avatar-menu svg');
	}

	get avatarMenu(): AvatarMenu {
		return new AvatarMenu(
			this.locator.locator('div[role="tooltip"]:has(span[data-testid="account_name"])')
		);
	}
}

export class AvatarMenu {
	constructor(readonly locator: Locator) {}

	get accountName(): Locator {
		return this.locator.locator('span[data-testid="account_name"]');
	}
	get accountEmail(): Locator {
		return this.locator.locator('span[data-testid="account_email"]');
	}

	get signoutButton(): Locator {
		return this.locator.locator('button:text("ログアウト")');
	}
}
