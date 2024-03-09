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
	idToken: string;
	channelId: bigint;
	limit?: number;
	after?: bigint;
}): Promise<ChatMessage[]> => {
	const { idToken, channelId, after } = arg;
	let { limit } = arg;
	limit = limit ?? 50;
	const { response } = await client.list({ idToken, channelId, limit, after });
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
	idToken: string;
	id: bigint;
}): Promise<ChatMessage> => {
	const { idToken, id } = arg;
	const { response } = await client.show({ idToken, id });
	return response;
};

export const createChatMessage = async (arg: {
	idToken: string;
	channelId: bigint;
	content: string;
}): Promise<{ id: string }> => {
	const { idToken, channelId, content } = arg;
	const { response } = await client.create({ idToken, channelId, content });
	return { id: response.id.toString() };
};

export const updateChatMessage = async (arg: {
	idToken: string;
	id: bigint;
	content: string;
}): Promise<void> => {
	const { idToken, id, content } = arg;
	if (!id) return;
	await client.update({ idToken, id, content });
};

export const deleteChatMessage = async (arg: { idToken: string; id: bigint }): Promise<void> => {
	if (!arg.id) return;
	const { idToken, id } = arg;
	await client.delete({ idToken, id });
};
