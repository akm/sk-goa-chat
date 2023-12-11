import type { Channel } from '$lib/models/channel';
import { listChannels } from '$lib/server/grpc_channels_client';
// import type { ServerLoadEvent } from '@sveltejs/kit';

export async function load(): Promise<{ channels: Channel[] }> {
	const channels = await listChannels();
	return { channels };
}
