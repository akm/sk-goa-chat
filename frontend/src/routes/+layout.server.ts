import type { LayoutServerLoad } from './$types';
import { listChannels } from '$lib/server/grpc_channels_client';

// get `locals.user` and pass it to the `page` store
export const load: LayoutServerLoad = async ({ locals }) => {
	const channels = locals.idToken ? await listChannels({ idToken: locals.idToken }) : [];
	return { user: locals.user, channels };
};
