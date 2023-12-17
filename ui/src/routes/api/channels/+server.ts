import { error, json, type RequestEvent } from '@sveltejs/kit';
import { createChannel } from '$lib/server/grpc_channels_client';

// See https://kit.svelte.jp/docs/types#public-types-requestevent
//     https://kit.svelte.jp/docs/routing#server-receiving-data
//     https://learn.svelte.jp/tutorial/event
export async function POST(event: RequestEvent) {
	const { request } = event;
	const data = await request.json();
	const { name } = data;
	if (!name) throw error(400, 'name is required');
	const { id } = await createChannel({ name: name.toString() });
	return json({ id, name }, { status: 201 });
}
