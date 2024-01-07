import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

const apisvrOrigin = process.env.APISVR_ORIGIN ?? 'http://localhost:8000';

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	// https://vitejs.dev/config/server-options.html#server-proxy
	server: {
		proxy: {
			'/api': {
				target: apisvrOrigin,
				changeOrigin: false,
				ws: false
			}
		}
	}
});
