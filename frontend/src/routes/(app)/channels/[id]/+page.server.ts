import type { Channel } from '$lib/models/channel';
import type { ChatMessage } from '$lib/models/chat_message';
import { showChannel } from '$lib/server/grpc_channels_client';
import { listChatMessages } from '$lib/server/grpc_chat_messages_client';
import { redirect } from '@sveltejs/kit';
import type { ServerLoadEvent } from '@sveltejs/kit';

export async function load(
	event: ServerLoadEvent
): Promise<{ channel: Channel; messages: ChatMessage[] }> {
	if (!event.params.id) {
		throw redirect(304, '/');
	}
	if (!event.locals.idToken) {
		throw redirect(304, '/');
	}
	const channelID = BigInt(event.params.id);
	if (!channelID) {
		throw redirect(304, '/');
	}
	const channel = await showChannel({ idToken: event.locals.idToken, id: channelID });
	// return { channel }; // Error: Data returned from `load` while rendering /channels/[channel_id] is not serializable: Cannot stringify arbitrary non-POJOs (data.channel) というエラーが。

	const messages = await listChatMessages({
		idToken: event.locals.idToken,
		channelId: channelID
	});

	const lastMessageId = messages.length == 0 ? 0 : messages[messages.length - 1].id;
	return { channel: { id: channel.id, name: channel.name }, messages, lastMessageId };
}
