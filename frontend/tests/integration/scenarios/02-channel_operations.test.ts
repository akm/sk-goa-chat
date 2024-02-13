import { expect, test } from '@playwright/test';
import { loadCookiesFrom } from '../steps/auth';
import { ChannelListPane } from '../pom/panes/channel_list';
import { foo } from './config';

test('operate channels', async ({ page, context }) => {
	await loadCookiesFrom(context, foo.cookieFile);

	await test.step('check to be signed in', async () => {
		await page.goto('/');
		const channelList = new ChannelListPane(page);
		await expect(channelList.list.itemByName('general')).toBeVisible();
	});
});
