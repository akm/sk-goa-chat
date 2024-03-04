import { defineConfig, devices } from '@playwright/test';

export default defineConfig({
	// timeout: 5 * 60_000,
	use: {
		headless: process.env.HEADED != 'true',
		launchOptions: {
			slowMo: Number(process.env.SLOMO || 0) // テスト実行時のスローモーション。デモなら 2000 くらいがよい
		},
		baseURL: 'http://localhost:' + (process.env.APP_UISVR_HTTP_PORT ?? '4173'),
		trace: 'retain-on-failure',

		// https://playwright.dev/docs/api/class-testoptions#test-options-video
		video: process.env.VIDEO === 'on' ? 'on' : 'off'
	},

	// https://playwright.dev/docs/api/class-testconfig#test-config-web-server
	webServer: [
		{
			command: `make -C tests/integration apisvr_run`,
			port: Number(process.env.APP_APISVR_HTTP_PORT ?? 4173),
			reuseExistingServer: !process.env.CI,
			stdout: 'pipe'
		},
		{
			command: `make -C tests/integration uisvr_run`,
			port: Number(process.env.APP_UISVR_HTTP_PORT ?? 4173),
			reuseExistingServer: !process.env.CI,
			stdout: 'pipe'
		}
	],

	testDir: 'tests/integration/scenarios',
	testMatch: /(.+\.)?test\.ts/,
	workers: 1,

	// https://playwright.dev/docs/test-reporters#github-actions-annotations
	reporter: process.env.CI ? 'github' : 'list',

	// https://playwright.dev/docs/browsers
	// https://playwright.dev/docs/api/class-testconfig#test-config-projects
	projects: [
		/* Test against desktop browsers */
		{
			name: 'chromium',
			use: { ...devices['Desktop Chrome'] }
		},
		{
			name: 'firefox',
			use: { ...devices['Desktop Firefox'] }
		},
		{
			name: 'webkit',
			use: { ...devices['Desktop Safari'] }
		},
		/* Test against mobile viewports. */
		{
			name: 'Mobile Chrome',
			use: { ...devices['Pixel 7'] }
		},
		{
			name: 'Mobile Safari',
			use: { ...devices['iPhone 14'] }
		},
		/* Test against branded browsers. */
		{
			name: 'Google Chrome',
			use: { ...devices['Desktop Chrome'], channel: 'chrome' }
		},
		{
			name: 'Apple Safari',
			use: { ...devices['Desktop Safari'], channel: 'webkit' }
		},
		{
			name: 'Microsoft Edge',
			use: { ...devices['Desktop Edge'], channel: 'msedge' }
		}
	]
});
