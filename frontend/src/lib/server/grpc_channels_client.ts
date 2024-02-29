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

export const listChannels = async (arg: { idToken: string }): Promise<Channel[]> => {
	const { idToken } = arg;
	const { response } = await client.list({ idToken });
	return (
		response.items?.field.map((item) => {
			return {
				id: item.id,
				name: item.name
			};
		}) ?? []
	);
};

export const showChannel = async (arg: { idToken: string; id: bigint }): Promise<Channel> => {
	const { idToken, id } = arg;
	const { response } = await client.show({ idToken, id });
	return response;
};

export const createChannel = async (arg: {
	idToken: string;
	name: string;
}): Promise<{ id: string }> => {
	const { idToken, name } = arg;
	const { response } = await client.create({ idToken, name });
	return { id: response.id.toString() };
};

export const updateChannel = async (arg: {
	idToken: string;
	id: bigint;
	name: string;
}): Promise<void> => {
	const { idToken, id, name } = arg;
	if (!id) return;
	await client.update({ idToken, id, name });
};

export const deleteChannel = async (arg: { idToken: string; id: bigint }): Promise<void> => {
	if (!arg.id) return;
	const { idToken, id } = arg;
	await client.delete({ idToken, id });
};
