import type { PlaywrightTestConfig } from '@playwright/test';

const config: PlaywrightTestConfig = {
	// timeout: 5 * 60_000,
	use: {
		headless: process.env.HEADED != 'true',
		launchOptions: {
			slowMo: Number(process.env.SLOMO || 0) // テスト実行時のスローモーション。デモなら 2000 くらいがよい
		},
		baseURL: 'http://localhost:' + (process.env.APP_UISVR_HTTP_PORT ?? '4173'),
		trace: 'retain-on-failure'
	},

	// https://playwright.dev/docs/api/class-testconfig#test-config-web-server
	webServer: [
		{
			command: 'make -C tests/integration apisvr_run',
			port: Number(process.env.APP_APISVR_HTTP_PORT ?? 4173),
			reuseExistingServer: !process.env.CI
		},
		{
			command: 'make -C tests/integration uisvr_run',
			port: Number(process.env.APP_UISVR_HTTP_PORT ?? 4173),
			reuseExistingServer: !process.env.CI
		}
	],

	testDir: 'tests/integration/scenarios',
	testMatch: /(.+\.)?test\.ts/,
	workers: 1
};

export default config;
