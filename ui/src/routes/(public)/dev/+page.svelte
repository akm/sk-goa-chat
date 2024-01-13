<script lang="ts">
	import {
		Button,
		Listgroup,
		Heading,
		Avatar,
		NavHamburger,
		Dropdown,
		DropdownHeader,
		DropdownItem,
		Modal,
		Label,
		Input,
		Hr
	} from 'flowbite-svelte';
	import { CogOutline, TrashBinSolid, InfoCircleSolid } from 'flowbite-svelte-icons';

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
	<div class="flex flex-auto flex-shrink-0 flex-col rounded-2xl p-4">
		<div class="flex flex-row mb-4">
			<div class="flex-none w-12">
				<img src="/logo192.png" class="mr-3 h-6 sm:h-9" alt="SK Goa Chat Logo" />
			</div>
			<div class="flex-shrink pt-1">
				<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">
					SK Goa Chat
				</span>
			</div>
		</div>
	</div>

	<div class="mb-4 flex h-full flex-col overflow-x-auto">
		<div class="flex h-full flex-col">
			<div class="flex-row">
				{#if channelLinks.length > 0}
					<Listgroup active items={channelLinks} let:item class="w-48" data-testid="channel_list">
						{item.name}
					</Listgroup>
				{/if}

				<Button class="mt-4" color="alternative" href="/channels/new">New Channel</Button>
			</div>
		</div>
	</div>
	<div class="flex w-full flex-row items-center rounded-xl bg-white">
		<div class="flex">
			<div class="flex flex-row">
				<div class="flex items-center space-x-4 rtl:space-x-reverse">
					<Avatar id="avatar-menu" rounded />
					<div class="space-y-1 font-medium dark:text-white">
						<div>Jese Leos</div>
						<div class="text-sm text-gray-500 dark:text-gray-400">Joined in August 2014</div>
					</div>
				</div>

				<Dropdown placement="bottom" triggeredBy="#avatar-menu">
					<DropdownItem>ログアウト</DropdownItem>
				</Dropdown>
			</div>
		</div>
	</div>
</div>

<div class="flex h-full flex-auto flex-col p-6">
	<div class="flex h-full flex-auto flex-shrink-0 flex-col rounded-2xl bg-gray-100 p-4">
		<div class="flex flex-row">
			<Heading tag="h3" class="mb-4">
				general
				<Button pill={true} outline={true} class="!p-2">
					<CogOutline class="w-4 h-4 text-primary-700" />
				</Button>
			</Heading>
		</div>

		<div class="mb-4 flex h-full flex-col overflow-x-auto">
			<div class="flex h-full flex-col">
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
		<div class="flex h-32 w-full flex-row items-center rounded-xl bg-white">
			<div class="flex-grow">
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
