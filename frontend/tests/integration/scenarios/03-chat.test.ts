import { test, expect } from '@playwright/test';

import { loadIndexedDBDataFrom } from '../steps/indexeddb';
import { foo, bar, firebaseIndexedDBConfig } from './config';

import { signup } from '../steps/signup';
import { ChannelListPane } from '../pom/panes/channel_list';
import { ChatPage } from '../pom/pages/chat_page';

test('show signin page when go to root', async ({ page, browser }) => {
	await page.pause();

	const { dbName, objectStoreName, keyPath } = firebaseIndexedDBConfig;
	await loadIndexedDBDataFrom(page, dbName, objectStoreName, keyPath, foo.credentialFilePath);
	const channelList = new ChannelListPane(page);

	const chatPage1 = new ChatPage(page);
	await test.step('デフォルトのチャンネルの確認', async () => {
		await page.goto('/');
		await expect(channelList.list.itemByName('general')).toBeVisible();
		await expect(channelList.list.itemByName('random')).toBeVisible();
		await expect(chatPage1.title('general')).toBeVisible();
	});

	const context2 = await browser.newContext();
	const page2 = await context2.newPage();
	await signup(page2, bar);
	const chatPage2 = new ChatPage(page2);
	await test.step('デフォルトのチャンネルの確認', async () => {
		await page2.goto('/');
		await expect(channelList.list.itemByName('general')).toBeVisible();
		await expect(channelList.list.itemByName('random')).toBeVisible();
		await expect(chatPage2.title('general')).toBeVisible();
	});

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
