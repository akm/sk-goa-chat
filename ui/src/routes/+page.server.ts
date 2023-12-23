import { redirect } from '@sveltejs/kit';
import type { ServerLoadEvent } from '@sveltejs/kit';

export async function load(event: ServerLoadEvent): Promise<void> {
	if (!event.locals.user) {
		throw redirect(302, '/signin');
	}
}
