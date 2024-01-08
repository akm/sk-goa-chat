import { ChannelCredentials } from '@grpc/grpc-js';
import { GrpcTransport } from '@protobuf-ts/grpc-transport';

import type { Channel } from '$lib/models/channel';
import { ChannelsClient, type IChannelsClient } from './protos/goagen_services_channels.client';

const apisvrGrpcPort = import.meta.env.VITE_APISVR_GRPC_PORT ?? '8080';

const transport = new GrpcTransport({
	host: 'localhost:' + apisvrGrpcPort,
	channelCredentials: ChannelCredentials.createInsecure()
});

const client: IChannelsClient = new ChannelsClient(transport);

export const listChannels = async (arg: { sessionId: string }): Promise<Channel[]> => {
	const { sessionId } = arg;
	const { response } = await client.list({ sessionId });
	return (
		response.items?.field.map((item) => {
			return {
				id: item.id,
				name: item.name
			};
		}) ?? []
	);
};

export const showChannel = async (arg: { sessionId: string; id: bigint }): Promise<Channel> => {
	const { sessionId, id } = arg;
	const { response } = await client.show({ sessionId, id });
	return response;
};

export const createChannel = async (arg: {
	sessionId: string;
	name: string;
}): Promise<{ id: string }> => {
	const { sessionId, name } = arg;
	const { response } = await client.create({ sessionId, name });
	return { id: response.id.toString() };
};

export const updateChannel = async (arg: {
	sessionId: string;
	id: bigint;
	name: string;
}): Promise<void> => {
	const { sessionId, id, name } = arg;
	if (!id) return;
	await client.update({ sessionId, id, name });
};

export const deleteChannel = async (arg: { sessionId: string; id: bigint }): Promise<void> => {
	if (!arg.id) return;
	const { sessionId, id } = arg;
	await client.delete({ sessionId, id });
};
