import { redirect } from '@sveltejs/kit';
import type { ServerLoadEvent } from '@sveltejs/kit';
import { listChannels } from '$lib/server/grpc_channels_client';

export async function load(event: ServerLoadEvent): Promise<void> {
	if (!event.locals.user) {
		throw redirect(302, '/signin');
	}

	const channels = event.locals.uid ? await listChannels({ uid: event.locals.uid }) : undefined;
	if (channels && channels.length > 0) {
		throw redirect(302, `/channels/${channels[0].id}`);
	}
}
