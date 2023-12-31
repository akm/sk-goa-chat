import { expect, test } from '@playwright/test';

test('show signin page when go to root', async ({ page }) => {
	await page.goto('/');
	// await page.pause();
	await expect(page).toHaveURL(/signin/);
	await expect(page.locator('nav span:text("SK Goa Chat")')).toBeVisible();
	await expect(page.locator('h3:text("ログイン")')).toBeVisible();
});
