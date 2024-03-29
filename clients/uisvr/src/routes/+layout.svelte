<script lang="ts">
	import { Button, Avatar, Dropdown, DropdownItem, Listgroup } from 'flowbite-svelte';

	import '../app.pcss';
	import { page } from '$app/stores';
	import { auth, waitUntilSignedOut } from '$lib/firebase/auth';
	import { closeWebSockets } from '$lib/websockets';
	import type { Channel } from '$lib/models/channel';

	const signout = async () => {
		try {
			closeWebSockets();
		} catch (e) {
			console.log('failed to close websockets: ', e);
		}
		console.log('+layout.svelte signOut calling');
		try {
			// https://firebase.google.com/docs/reference/js/auth.auth.md#authsignout
			await auth.signOut();
			console.log('+layout.svelte signOut called');
			localStorage.removeItem('idToken');
			localStorage.removeItem('refreshToken');
		} catch (e) {
			console.log('failed to signOut: ', e);
		}
		await waitUntilSignedOut();
		await new Promise((resolve) => setTimeout(resolve, 1_000));

		console.log('+layout.svelte setting window.location.href');
		window.location.href = $page.url.origin + '/signin';
	};

	// user の例: {id: 'nUhxKTpuXq4phNaBp1NF6Vp605wJ', name: 'Foo', email: 'foo@example.com'}
	const user = $page.data.user;
	console.log('+layout.svelte script user:', user);

	const channelLinks = $page.data.channels
		? ($page.data.channels as Channel[]).map((channel) => ({
				name: channel.name,
				href: `/channels/${channel.id}`
		  }))
		: undefined;
</script>

<div class="flex h-full w-full flex-row overflow-x-hidden">
	<div class="flex w-64 flex-shrink-0 flex-col bg-white py-2 pl-6 pr-2" data-testid="sidebar">
		<div class="flex flex-auto flex-shrink-0 flex-col rounded-2xl p-4">
			<div class="flex flex-row mb-4">
				<div class="flex-none w-12">
					<a href="/">
						<img src="/logo192.png" class="mr-3 h-6 sm:h-9" alt="SK Goa Chat Logo" />
					</a>
				</div>
				<div class="flex-shrink pt-1">
					<a href="/">
						<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">
							SK Goa Chat
						</span>
					</a>
				</div>
			</div>
		</div>

		<div class="mb-4 flex h-full flex-col overflow-x-auto">
			<div class="flex h-full flex-col">
				<div class="flex-row" data-testid="channel_list_pane">
					{#if channelLinks && channelLinks.length > 0}
						<Listgroup active items={channelLinks} let:item class="w-48" data-testid="channel_list">
							{item.name}
						</Listgroup>

						<Button class="mt-4" color="alternative" href="/channels/new">New Channel</Button>
					{/if}
				</div>
			</div>
		</div>

		{#if user}
			<div class="flex w-full flex-row items-center rounded-xl bg-white">
				<div class="flex">
					<div class="flex flex-row">
						<div class="flex items-center space-x-4 rtl:space-x-reverse">
							<Avatar id="avatar-menu" rounded />
							<div class="space-y-1 font-medium dark:text-white">
								<div data-testid="account_name">{user.name}</div>
								<div data-testid="account_email" class="text-sm text-gray-500 dark:text-gray-400">
									{user.email}
								</div>
							</div>
						</div>

						<Dropdown placement="bottom" triggeredBy="#avatar-menu">
							<DropdownItem on:click={signout}>ログアウト</DropdownItem>
						</Dropdown>
					</div>
				</div>
			</div>
		{/if}
	</div>

	<slot />
</div>
