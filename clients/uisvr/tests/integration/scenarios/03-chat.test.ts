import { test, expect } from '@playwright/test';

import { getFoo, getBar } from './config';

import { signup } from '../steps/signup';
import { ChannelListPane } from '../pom/panes/channel_list';
import { NewChannelPage } from '../pom/pages/new_channel_page';
import { ChatPage } from '../pom/pages/chat_page';
import { signin } from '../steps/signin';

test('show signin page when go to root', async ({ page, browser }, workers) => {
	const foo = getFoo(workers.workerIndex);
	const bar = getBar(workers.workerIndex);

	await page.goto('/');
	await signin(page, foo);

	const channnelName = `channel-${workers.workerIndex}`;
	const newChannelPage = new NewChannelPage(page);
	await test.step('チャットテスト用チャンネルの追加', async () => {
		const chatPage = new ChatPage(page);
		const channelList = new ChannelListPane(page);
		await channelList.button('New Channel').click();
		await expect(page).toHaveURL(/channels\/new$/);
		await expect(newChannelPage.title).toBeVisible();
		await newChannelPage.inputField('Name').fill(channnelName);
		await newChannelPage.button('Create').click();
		await expect(chatPage.title(channnelName)).toBeVisible();
		await expect(channelList.list.itemByName(channnelName)).toBeVisible();

		await expect(channelList.list.itemByName('general')).toBeVisible();
		await expect(channelList.list.itemByName('random')).toBeVisible();
		await expect(channelList.list.itemByName(channnelName)).toBeVisible();
		await expect(chatPage.title(channnelName)).toBeVisible();
	});

	const context2 = await browser.newContext();
	const page2 = await context2.newPage();
	await signup(page2, bar);
	const chatPage2 = new ChatPage(page2);
	await test.step('デフォルトのチャンネルの確認', async () => {
		await page2.goto('/');

		await page.waitForTimeout(1_000); // TODO このスリープを除去

		const channelList = new ChannelListPane(page2);
		await expect(channelList.list.itemByName('general')).toBeVisible();
		await expect(channelList.list.itemByName('random')).toBeVisible();
		await expect(channelList.list.itemByName(channnelName)).toBeVisible();
		await expect(chatPage2.title('general')).toBeVisible();
		await channelList.list.itemByName(channnelName).click();
		await expect(chatPage2.title(channnelName)).toBeVisible();
	});

	const chatPage1 = new ChatPage(page);

	await test.step('foo sends messages', async () => {
		const fooMsg1 = 'foo message 1';
		await chatPage1.textarea.fill(fooMsg1);
		await chatPage1.button('Send').click();
		await expect(chatPage1.messagePane.locator(`p:text("${fooMsg1}")`)).toBeVisible();
		await expect(chatPage2.messagePane.locator(`p:text("${fooMsg1}")`)).toBeVisible();
	});

	await test.step('bar sends messages', async () => {
		const barMsg1 = 'bar message 1';
		await chatPage2.textarea.fill(barMsg1);
		await chatPage2.button('Send').click();
		await expect(chatPage1.messagePane.locator(`p:text("${barMsg1}")`)).toBeVisible();
		await expect(chatPage2.messagePane.locator(`p:text("${barMsg1}")`)).toBeVisible();
	});
});
