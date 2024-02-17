import fs from 'fs';
import { test } from '@playwright/test';
import type { BrowserContext } from '@playwright/test';

export const authFilepath = (email: string): string => {
	return `tmp/cookies-${email.split('@')[0]}.json`;
};

export const saveCookiesTo = async (context: BrowserContext, filePath: string): Promise<void> => {
	await test.step(`save cookies to ${filePath}`, async () => {
		const cookies = await context.cookies();
		fs.writeFileSync(filePath, JSON.stringify(cookies), 'utf-8');
	});
};

export const loadCookiesFrom = async (context: BrowserContext, filePath: string): Promise<void> => {
	await test.step(`load cookies from ${filePath}`, async () => {
		const cookies = JSON.parse(fs.readFileSync(filePath, 'utf-8'));
		await context.clearCookies();
		await context.addCookies(cookies);
	});
};
