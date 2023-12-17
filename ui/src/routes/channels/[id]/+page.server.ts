import type { Channel } from '$lib/models/channel';
import { showChannel } from '$lib/server/grpc_channels_client';
import { redirect } from '@sveltejs/kit';
import type { ServerLoadEvent } from '@sveltejs/kit';

export async function load(event: ServerLoadEvent): Promise<{ channel: Channel }> {
	const channelID = BigInt(event.params.id || '');
	if (!channelID) {
		throw redirect(304, '/');
	}
	const channel = await showChannel({ id: channelID });
	// return { channel }; // Error: Data returned from `load` while rendering /channels/[channel_id] is not serializable: Cannot stringify arbitrary non-POJOs (data.channel) というエラーが。
	return { channel: { id: channel.id, name: channel.name } };
}
