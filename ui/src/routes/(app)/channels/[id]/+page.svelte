<script lang="ts">
	import { Heading, TextPlaceholder, Alert } from 'flowbite-svelte';
	import { Button, Modal, Label, Input, Hr } from 'flowbite-svelte';
	import { CogOutline, TrashBinSolid, InfoCircleSolid } from 'flowbite-svelte-icons';

	import type { Channel } from '$lib/models/channel';
	import type { ChatMessage } from '$lib/models/chat_message';
	import { channelOptionsEqual } from '@grpc/grpc-js/build/src/channel-options.js';
	import { notificationsSocket } from '$lib/websockets';
	import { onDestroy, onMount } from 'svelte';

	export let data: {
		channel: Channel;
		messages: ChatMessage[];
		lastMessageId: number;
	};

	let name = data.channel.name;
	let errorMessage = '';

	let settingVisible = false;

	const nofiticationHandler = (event) => {
		// console.log('event', event);
		// console.log('event.data', event.data);
		const notification = JSON.parse(event.data);
		// console.log('notification', notification);
		const chID = Number(data.channel.id);
		if (notification.channel_ids.includes(chID)) {
			readLaterMessages()
				.then(() => {
					console.log('update by notification OK');
				})
				.catch((err) => {
					console.log('update by notification failed', err);
				});
		} else {
			console.log('notification does not include channel id', {
				channel_ids: notification.channel_ids,
				channel_id: chID
			});
		}
	};

	let ws: WebSocket;
	onMount(() => {
		ws = notificationsSocket();
		ws.addEventListener('message', nofiticationHandler);
		scrollToLatestChat('instant');
	});
	onDestroy(() => {
		if (ws) ws.removeEventListener('message', nofiticationHandler);
	});

	const updateChannel = async () => {
		const result = await fetch(`/api/channels/${data.channel.id}`, {
			method: 'PUT',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ name })
		});
		const json = await result.json();
		console.log('json', json);
		if (!json.id) {
			errorMessage = json.message;
			return;
		}
		window.location.reload();
	};

	const deleteChannel = async () => {
		const result = await fetch(`/api/channels/${data.channel.id}`, {
			method: 'DELETE',
			headers: { 'Content-Type': 'application/json' }
		});
		const json = await result.json();
		console.log('json', json);
		if (!json.id) {
			errorMessage = json.message;
			return;
		}
		window.location.href = '/';
	};

	let textarea: HTMLTextAreaElement;
	const postMessage = async () => {
		const result = await fetch(`/api/chat_messages`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ channel_id: Number(data.channel.id), content: textarea.value })
		});
		const json = await result.json();
		console.log('json', json);
		if (!json.id) {
			errorMessage = json.message;
			return;
		}
		textarea.value = '';
		textarea.focus();
	};

	const readNewMessages = async (reqPath: string): Promise<ChatMessage[]> => {
		const result = await fetch(reqPath);
		const json = await result.json();
		console.log('json', json);
		if (!json.items) {
			errorMessage = json.message;
			throw json.message;
		}
		return json.items.map((msg) => ({
			id: msg.id,
			createdAt: msg.created_at,
			updatedAt: msg.updated_at,
			channelId: msg.channel_id,
			userId: msg.user_id,
			userName: msg.user_name,
			content: msg.content
		}));
	};

	const uniqSort = (array: ChatMessage[]) => {
		return array
			.filter((item, index) => {
				return array.findIndex((item2) => item.id === item2.id) === index;
			})
			.sort((a, b) => {
				if (a.id < b.id) {
					return -1;
				}
				if (a.id > b.id) {
					return 1;
				}
				return 0;
			});
	};

	let scrollContainer: HTMLDivElement;
	const latestChatVisible = (): boolean => {
		const pos =
			scrollContainer.scrollHeight - (scrollContainer.scrollTop + scrollContainer.clientHeight);
		return pos < 100;
	};
	const scrollToLatestChat = (behavior?: ScrollBehavior) => {
		scrollContainer.scrollTo({
			top: scrollContainer.scrollHeight - scrollContainer.clientHeight - 50,
			behavior: behavior
		});
	};

	const readLaterMessages = async () => {
		const latestVisible = latestChatVisible();
		const newMessages = await readNewMessages(
			`/api/chat_messages?channel_id=${data.channel.id}&after=${data.lastMessageId}&limit=50`
		);
		data.messages = uniqSort([...data.messages, ...newMessages]);
		data.lastMessageId = Number(data.messages[data.messages.length - 1].id);
		if (latestVisible) {
			setTimeout(() => {
				scrollToLatestChat('smooth');
			}, 10);
		}
	};

	const readEarlierMessages = async () => {
		const earliestId = data.messages[0].id;
		const newMessages = await readNewMessages(
			`/api/chat_messages?channel_id=${data.channel.id}&before=${earliestId}&limit=50`
		);
		data.messages = uniqSort([...newMessages, ...data.messages]);
	};
</script>

<div class="flex h-full flex-auto flex-col p-6">
	<div class="flex h-full flex-auto flex-shrink-0 flex-col rounded-2xl bg-gray-100 p-4">
		<div class="flex flex-row">
			<Heading tag="h3" class="mb-4">
				{data.channel.name}
				<Button pill={true} outline={true} class="!p-2" on:click={() => (settingVisible = true)}>
					<CogOutline class="w-4 h-4 text-primary-700" />
				</Button>
			</Heading>
		</div>

		<div class="mb-4 flex h-full flex-col overflow-x-auto" bind:this={scrollContainer}>
			<div class="flex h-full flex-col">
				<div class="flex justify-center">
					<Button on:click={readEarlierMessages} class="mt-4" color="alternative"
						>Read previous messages</Button
					>
				</div>

				{#each data.messages as msg (msg.id)}
					<div class="flex items-start mb-4">
						<!-- <div class="flex-shrink-0">
		<img
			class="w-10 h-10 rounded-full"
			src={msg.user.avatarUrl}
			alt={msg.user.name}
		/>
	</div> -->
						<div class="flex flex-col flex-1 w-0 ms-3">
							<div class="flex items-center justify-between">
								<h4 class="text-sm font-medium text-gray-900 dark:text-white">
									{msg.userName}
								</h4>
								<p class="text-sm text-gray-500">
									{msg.createdAt}
								</p>
							</div>
							<p class="mt-1 text-sm text-gray-700 dark:text-white">
								{msg.content}
							</p>
						</div>
					</div>
				{/each}

				<div class="flex justify-center">
					<Button on:click={readLaterMessages} class="mt-4" color="alternative"
						>Read next messages</Button
					>
				</div>
			</div>
		</div>
		<div class="flex h-32 w-full flex-row items-center rounded-xl bg-white">
			<div class="flex-grow">
				<div class="relative w-full">
					<div class="flex flex-row">
						<textarea
							bind:this={textarea}
							on:keypress={(e) => {
								if (e.key === 'Enter' && (e.altKey || e.metaKey || e.ctrlKey)) {
									// Chrome では metaKey は反応しない
									e.preventDefault();
									postMessage();
								}
							}}
							class="grow h-24 p-2 border border-gray-300 rounded-md"
						/>
						<div class="flex-0 h-24">
							<Button class="m-2 h-20" on:click={postMessage}>Send</Button>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<Modal bind:open={settingVisible} size="xs" autoclose={false} class="w-full">
	<h3 class="text-xl font-medium text-gray-900 dark:text-white">Channel Settings</h3>

	{#if errorMessage}
		<Alert
			type="error"
			class="mb-4"
			dismissable
			on:close={() => {
				errorMessage = '';
			}}
		>
			<InfoCircleSolid slot="icon" class="w-4 h-4" />
			{errorMessage}
		</Alert>
	{/if}

	<div class="flex flex-col space-y-6">
		<Label class="space-y-2">
			<span>Name</span>
			<Input type="text" name="name" required bind:value={name} />
		</Label>
		<div class="flex">
			<Button type="submit" class="mr-4" on:click={updateChannel}>Update</Button>
			<Button
				color="alternative"
				on:click={() => {
					settingVisible = false;
				}}>Cancel</Button
			>
		</div>
	</div>

	<Hr />

	<h3 class="text-xl font-medium text-red-700 dark:text-white">Danger Done</h3>

	<Button on:click={deleteChannel}>
		<TrashBinSolid class="w-3.5 h-3.5 me-2" />
		Delete this channel forever
	</Button>
</Modal>
