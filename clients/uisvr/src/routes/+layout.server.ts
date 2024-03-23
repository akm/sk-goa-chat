import type { LayoutServerLoad } from './$types';
import { listChannels } from '$lib/server/grpc_channels_client';

// get `locals.user` and pass it to the `page` store
export const load: LayoutServerLoad = async ({ locals }) => {
	console.log('layout.server.ts load locals:', locals);
	const channels = locals.uid ? await listChannels({ uid: locals.uid }) : [];
	return { user: locals.user, channels };
};
