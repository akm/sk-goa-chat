import type { Channel } from '$lib/models/channel';
import { showChannel } from '$lib/server/grpc_channels_client';
import { redirect } from '@sveltejs/kit';
import type { ServerLoadEvent } from '@sveltejs/kit';

export async function load(event: ServerLoadEvent): Promise<{ channel: Channel }> {
	const channel_id = BigInt(event.params.channel_id || '');
	if (!channel_id) {
		throw redirect(304, '/');
	}
	const channel = await showChannel({ id: channel_id });
	// return { channel }; // Error: Data returned from `load` while rendering /channels/[channel_id] is not serializable: Cannot stringify arbitrary non-POJOs (data.channel) というエラーが。
	return { channel: { id: channel.id, name: channel.name } };
}
