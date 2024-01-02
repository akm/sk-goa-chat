import type { PlaywrightTestConfig } from '@playwright/test';

const config: PlaywrightTestConfig = {
	// timeout: 5 * 60_000,
	use: {
		headless: process.env.HEADED != 'true',
		launchOptions: {
			slowMo: Number(process.env.SLOMO || 0) // テスト実行時のスローモーション。デモなら 2000 くらいがよい
		},
		baseURL: 'http://localhost:4173',
		trace: 'retain-on-failure'
	},

	// https://playwright.dev/docs/api/class-testconfig#test-config-web-server
	webServer: [
		{
			command: 'make -C tests/integration apisvr_run',
			port: 8001,
			reuseExistingServer: !process.env.CI
		},
		{
			command: 'npm run build && npm run preview >> tests/integration/log/uisvr.log 2>&1',
			port: 4173,
			reuseExistingServer: !process.env.CI,
			env: {
				VITE_GOOGLE_CLOUD_PROJECT: 'sk-goa-chat',
				// https://firebase.google.com/docs/emulator-suite/connect_auth?hl=ja
				FIREBASE_AUTH_EMULATOR_HOST: '127.0.0.1:9090', // VITE_FIREBASE_AUTH_EMULATOR_HOST と重複していますが、この環境変数も必要です
				VITE_FIREBASE_AUTH_EMULATOR_HOST: '127.0.0.1:9090',
				// VITE_APISVR_HTTP_PORT: '8001',
				VITE_APISVR_GRPC_PORT: '8081'
			}
		}
	],

	testDir: 'tests',
	testMatch: /(.+\.)?(test|spec)\.[jt]s/
};

export default config;
