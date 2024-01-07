import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	// https://vitejs.dev/config/server-options.html#server-proxy
	server: {
		proxy: {
			'/api': {
				target: 'http://localhost:8000',
				changeOrigin: false,
				ws: false
			}
		}
	}
});
