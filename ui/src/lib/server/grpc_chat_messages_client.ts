import { ChannelCredentials } from '@grpc/grpc-js';
import { GrpcTransport } from '@protobuf-ts/grpc-transport';

import type { ChatMessage } from '$lib/models/chat_message';
import {
	ChatMessagesClient,
	type IChatMessagesClient
} from './protos/goagen_services_chat_messages.client';

const apisvrGrpcPort = import.meta.env.VITE_APISVR_GRPC_PORT ?? '8080';

const transport = new GrpcTransport({
	host: 'localhost:' + apisvrGrpcPort,
	channelCredentials: ChannelCredentials.createInsecure()
});

const client: IChatMessagesClient = new ChatMessagesClient(transport);

export const listChatMessages = async (arg: {
	sessionId: string;
	channelId: bigint;
	limit?: number;
	after?: bigint;
}): Promise<ChatMessage[]> => {
	const { sessionId, channelId, after } = arg;
	let { limit } = arg;
	limit = limit ?? 50;
	const { response } = await client.list({ sessionId, channelId, limit, after });
	return (
		response.items?.field.map((item) => {
			return {
				id: item.id,
				channelId: item.channelId,
				userId: item.userId,
				userName: item.userName,
				content: item.content,
				createdAt: item.createdAt,
				updatedAt: item.updatedAt
			};
		}) ?? []
	);
};

export const showChatMessage = async (arg: {
	sessionId: string;
	id: bigint;
}): Promise<ChatMessage> => {
	const { sessionId, id } = arg;
	const { response } = await client.show({ sessionId, id });
	return response;
};

export const createChatMessage = async (arg: {
	sessionId: string;
	channelId: bigint;
	content: string;
}): Promise<{ id: string }> => {
	const { sessionId, channelId, content } = arg;
	const { response } = await client.create({ sessionId, channelId, content });
	return { id: response.id.toString() };
};

export const updateChatMessage = async (arg: {
	sessionId: string;
	id: bigint;
	content: string;
}): Promise<void> => {
	const { sessionId, id, content } = arg;
	if (!id) return;
	await client.update({ sessionId, id, content });
};

export const deleteChatMessage = async (arg: { sessionId: string; id: bigint }): Promise<void> => {
	if (!arg.id) return;
	const { sessionId, id } = arg;
	await client.delete({ sessionId, id });
};