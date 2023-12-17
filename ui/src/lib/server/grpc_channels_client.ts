import { ChannelCredentials } from '@grpc/grpc-js';
import { GrpcTransport } from '@protobuf-ts/grpc-transport';

import type { Channel } from '$lib/models/channel';
import { ChannelsClient, type IChannelsClient } from './protos/goagen_services_channels.client';

const transport = new GrpcTransport({
	host: 'localhost:8080',
	channelCredentials: ChannelCredentials.createInsecure()
});

const client: IChannelsClient = new ChannelsClient(transport);

export const listChannels = async (): Promise<Channel[]> => {
	const { response } = await client.list({});
	return (
		response.items?.field.map((item) => {
			return {
				id: item.id,
				name: item.name
			};
		}) ?? []
	);
};

export const createChannel = async (arg: { name: string }): Promise<{ id: string }> => {
	const { name } = arg;
	const { response } = await client.create({ name: name });
	return { id: response.id.toString() };
};

export const updateChannel = async (arg: { id: bigint; name: boolean }): Promise<void> => {
	if (!arg.id) return;
	const { response } = await client.show({ id: BigInt(arg.id) });
	await client.update({
		id: BigInt(arg.id),
		name: response.name
	});
};

export const deleteChannel = async (arg: { id: bigint }): Promise<void> => {
	if (!arg.id) return;
	await client.delete({ id: arg.id });
};