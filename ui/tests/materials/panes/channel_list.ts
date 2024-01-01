import type { Locator, Page } from '@playwright/test';

export class ChannelListPane {
	readonly locator: Locator;
	constructor(page: Page) {
		this.locator = page.locator('[data-testid="channel_list_pane"]');
	}

	get list(): ChannelList {
		return new ChannelList(this.locator.locator('[data-testid="channel_list"]'));
	}

	button(name: 'New Channel'): Locator {
		return this.locator.locator(`a:text("${name}")`);
	}
}

export class ChannelList {
	constructor(readonly locator: Locator) {}

	itemByName(name: string): Locator {
		return this.locator.locator(`a:text("${name}")`);
	}
}
