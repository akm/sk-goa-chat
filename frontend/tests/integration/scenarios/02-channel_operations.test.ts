import { expect, test } from '@playwright/test';
import { foo } from './config';

import { ChannelListPane } from '../pom/panes/channel_list';
import { NewChannelPage } from '../pom/pages/new_channel_page';
import { ChatPage } from '../pom/pages/chat_page';
import { ChannelSettingDialog } from '../pom/dialogs/channel_setting_dialog';
import { signin } from '../steps/signin';

test('operate channels', async ({ page }) => {
	const channelList = new ChannelListPane(page);

	await page.goto('/');
	await signin(page, foo);

	await test.step('デフォルトのチャンネルの確認', async () => {
		await page.goto('/');
		await expect(channelList.list.itemByName('general')).toBeVisible();
		await expect(channelList.list.itemByName('random')).toBeVisible();
	});

	const channel1 = 'new channel1';
	const newChannelPage = new NewChannelPage(page);
	const chatPage = new ChatPage(page);
	await test.step('チャンネルの追加', async () => {
		await channelList.button('New Channel').click();
		await expect(page).toHaveURL(/channels\/new$/);
		await expect(newChannelPage.title).toBeVisible();
		await newChannelPage.inputField('Name').fill(channel1);
		await newChannelPage.button('Create').click();
		await expect(chatPage.title(channel1)).toBeVisible();
		await expect(channelList.list.itemByName(channel1)).toBeVisible();
	});

	const newChannelName = 'new channel2';
	const channelSettingDialog = new ChannelSettingDialog(page);
	await test.step('チャンネルの名前の変更', async () => {
		await chatPage.configButton.click();
		await expect(channelSettingDialog.title).toBeVisible();
		await channelSettingDialog.inputField('Name').fill(newChannelName);
		await channelSettingDialog.button('Update').click();
		await expect(channelSettingDialog.title).not.toBeVisible();
		await expect(chatPage.title(newChannelName)).toBeVisible();
		await expect(channelList.list.itemByName(newChannelName)).toBeVisible();
	});

	await test.step('チャンネルの削除', async () => {
		await chatPage.configButton.click();
		await expect(channelSettingDialog.title).toBeVisible();
		await channelSettingDialog.button('Delete this channel forever').click();
		await expect(channelSettingDialog.title).not.toBeVisible();
		await expect(chatPage.title('general')).toBeVisible();
		await expect(channelList.list.itemByName(channel1)).not.toBeVisible();
		await expect(channelList.list.itemByName(newChannelName)).not.toBeVisible();
	});
});
