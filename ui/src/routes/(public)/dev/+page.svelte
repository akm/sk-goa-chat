<script lang="ts">
	import { Button, Listgroup, Modal, Label, Input, Hr } from 'flowbite-svelte';

	const channelLinks = [
		{ name: 'general', href: '/channels/1' },
		{ name: 'random', href: '/channels/2' }
	];

	const messages = [...Array(50)].map((_, i) => {
		return {
			id: i,
			userName: `user${(i % 4) + 1}`,
			createdAt: '2021-09-01 12:00:00',
			content: `message${i + 1}`
		};
	});

	let settingVisible = false;
</script>

<div class="flex w-64 flex-shrink-0 flex-col bg-white py-8 pl-6 pr-2">
	{#if channelLinks.length > 0}
		<Listgroup active items={channelLinks} let:item class="w-48" data-testid="channel_list">
			{item.name}
		</Listgroup>
	{/if}

	<Button class="mt-4" color="alternative" href="/channels/new">New Channel</Button>
</div>

<div class="flex h-full flex-auto flex-col p-6">
	<div class="flex h-full flex-auto flex-shrink-0 flex-col rounded-2xl bg-gray-100 p-4">
		<div class="mb-4 flex h-full flex-col overflow-x-auto">
			<div class="flex h-full flex-col">
				<div class="">
					<div class="flex flex-row justify-center">
						<Button class="mt-4" color="alternative">Read previous messages</Button>
					</div>

					{#each messages as msg (msg.id)}
						<div class="flex flex-row items-start mb-4">
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

					<div class="flex flex-row justify-center">
						<Button class="mt-4" color="alternative">Read next messages</Button>
					</div>
				</div>
			</div>
		</div>
		<div class="flex h-32 w-full flex-row items-center rounded-xl bg-white px-4">
			<div class="ml-4 flex-grow">
				<div class="relative w-full">
					<div class="flex flex-row">
						<textarea class="grow h-24 p-2 border border-gray-300 rounded-md" />
						<div class="flex-0 h-24">
							<Button class="m-2 h-20" on:click={postMessage}>Send</Button>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
