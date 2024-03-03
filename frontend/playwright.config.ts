import type { PlaywrightTestConfig } from '@playwright/test';

const config: PlaywrightTestConfig = {
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
			command: `make -C tests/integration apisvr_run >> ./tests/integration/log/playwright.config.log 2>&1`,
			port: Number(process.env.APP_APISVR_HTTP_PORT ?? 4173),
			reuseExistingServer: !process.env.CI
		},
		{
			command: `make -C tests/integration uisvr_run >> ./tests/integration/log/playwright.config.log 2>&1`,
			port: Number(process.env.APP_UISVR_HTTP_PORT ?? 4173),
			reuseExistingServer: !process.env.CI
		}
	],

	testDir: 'tests/integration/scenarios',
	testMatch: /(.+\.)?test\.ts/,
	workers: 1,

	// https://playwright.dev/docs/test-reporters#github-actions-annotations
	reporter: process.env.CI ? 'github' : 'list'
};

export default config;
