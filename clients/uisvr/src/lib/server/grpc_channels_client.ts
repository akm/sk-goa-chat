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

export const listChannels = async (arg: { uid: string }): Promise<Channel[]> => {
	const { uid } = arg;
	const { response } = await client.list({ uid });
	return (
		response.items?.field.map((item) => {
			return {
				id: item.id,
				name: item.name
			};
		}) ?? []
	);
};

export const showChannel = async (arg: { uid: string; id: bigint }): Promise<Channel> => {
	const { uid, id } = arg;
	const { response } = await client.show({ uid, id });
	return response;
};

export const createChannel = async (arg: {
	uid: string;
	name: string;
}): Promise<{ id: string }> => {
	const { uid, name } = arg;
	const { response } = await client.create({ uid, name });
	return { id: response.id.toString() };
};

export const updateChannel = async (arg: {
	uid: string;
	id: bigint;
	name: string;
}): Promise<void> => {
	const { uid, id, name } = arg;
	if (!id) return;
	await client.update({ uid, id, name });
};

export const deleteChannel = async (arg: { uid: string; id: bigint }): Promise<void> => {
	if (!arg.id) return;
	const { uid, id } = arg;
	await client.delete({ uid, id });
};
