import type { PlaywrightTestConfig } from '@playwright/test';

const config: PlaywrightTestConfig = {
	// https://playwright.dev/docs/api/class-testconfig#test-config-web-server
	webServer: [
		{
			command: 'make apisvr_run',
			port: 8001,
			reuseExistingServer: !process.env.CI
		},
		{
			command: 'npm run build && npm run preview',
			port: 4173,
			reuseExistingServer: !process.env.CI,
			env: {
				GOOGLE_CLOUD_PROJECT: 'sk-goa-chat',
				FIREBASE_AUTH_EMULATOR_HOST: '127.0.0.1:9090'
			}
		}
	],

	testDir: 'tests',
	testMatch: /(.+\.)?(test|spec)\.[jt]s/
};

export default config;
