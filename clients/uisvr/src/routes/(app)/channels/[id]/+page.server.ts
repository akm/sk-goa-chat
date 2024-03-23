import type { Channel } from '$lib/models/channel';
import type { ChatMessage } from '$lib/models/chat_message';
import { showChannel } from '$lib/server/grpc_channels_client';
import { listChatMessages } from '$lib/server/grpc_chat_messages_client';
import { redirect } from '@sveltejs/kit';
import type { ServerLoadEvent } from '@sveltejs/kit';

export async function load(
	event: ServerLoadEvent
): Promise<{ channel: Channel; messages: ChatMessage[]; lastMessageId: bigint | 0 }> {
	if (!event.params.id) {
		throw redirect(304, '/');
	}
	if (!event.locals.uid) {
		throw redirect(304, '/');
	}
	const channelID = BigInt(event.params.id);
	if (!channelID) {
		throw redirect(304, '/');
	}
	const channel = await showChannel({ uid: event.locals.uid, id: channelID });
	// return { channel }; // Error: Data returned from `load` while rendering /channels/[channel_id] is not serializable: Cannot stringify arbitrary non-POJOs (data.channel) というエラーが。

	const messages = await listChatMessages({
		uid: event.locals.uid,
		channelId: channelID
	});

	const lastMessageId = messages.length == 0 ? 0 : messages[messages.length - 1].id;
	return { channel: { id: channel.id, name: channel.name }, messages, lastMessageId };
}
