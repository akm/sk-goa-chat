import { json, error, type RequestEvent } from '@sveltejs/kit';
import { updateChannel, deleteChannel } from '$lib/server/grpc_channels_client';

// See https://kit.svelte.jp/docs/types#public-types-requestevent
//     https://kit.svelte.jp/docs/routing#server-receiving-data
//     https://learn.svelte.jp/tutorial/event
export async function PUT(event: RequestEvent) {
	const { params, request } = event;
	console.log(params);
	const id = BigInt(params.id || '');
	if (!id) throw error(404, 'id is not given');
	const data = await request.json();
	const { name } = data;
	console.log(data);
	if (!name) throw error(400, 'name is required');
	await updateChannel({ id, name: name.toString() });
	return json({ id: Number(id), name }, { status: 200 });
}

// See https://kit.svelte.jp/docs/types#public-types-requestevent
//     https://kit.svelte.jp/docs/routing#server-receiving-data
//     https://learn.svelte.jp/tutorial/event
export async function DELETE(event: RequestEvent) {
	const { params } = event;
	const id = BigInt(params.id || '');
	if (!id) throw error(404, 'id is not given');
	await deleteChannel({ id });
	return json({ id: Number(id) }, { status: 200 });
}
