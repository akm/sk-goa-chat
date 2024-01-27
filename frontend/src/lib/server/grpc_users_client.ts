import { ChannelCredentials } from '@grpc/grpc-js';
import { GrpcTransport } from '@protobuf-ts/grpc-transport';

import { UsersClient, type IUsersClient } from './protos/goagen_services_users.client';

const apisvrGrpcPort = import.meta.env.VITE_APISVR_GRPC_PORT ?? '8080';

const transport = new GrpcTransport({
	host: 'localhost:' + apisvrGrpcPort,
	channelCredentials: ChannelCredentials.createInsecure()
});

const client: IUsersClient = new UsersClient(transport);

export const createUser = async (arg: { email: string; name: string }): Promise<{ id: string }> => {
	const { email, name } = arg;
	const { response } = await client.create({ email, name });
	return { id: response.id.toString() };
};
