<script lang="ts">
	import {
		Navbar,
		NavBrand,
		NavHamburger,
		Avatar,
		Dropdown,
		DropdownItem,
		DropdownHeader
	} from 'flowbite-svelte';

	import '../app.pcss';
	import { page } from '$app/stores';
	import { deleteSession } from '$lib/session';

	const signout = async () => {
		await deleteSession();
		window.location.href = $page.url.origin + '/signin';
	};

	// user の例: {id: 'nUhxKTpuXq4phNaBp1NF6Vp605wJ', name: 'Foo', email: 'foo@example.com'}
	const user = $page.data.user;

	console.log('user', user);
</script>

<div class="mx-36 max-w-full my-2">
	<Navbar data-testid="header_nav">
		<NavBrand href="/">
			<img src="/logo192.png" class="mr-3 h-6 sm:h-9" alt="SK Goa Chat Logo" />
			<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
				>SK Goa Chat</span
			>
		</NavBrand>
		{#if user}
			<div class="flex items-center md:order-2">
				<Avatar id="avatar-menu" />
				<NavHamburger class1="w-full md:flex md:w-auto md:order-1" />
			</div>
			<Dropdown placement="bottom" triggeredBy="#avatar-menu">
				<DropdownHeader>
					<span data-testid="account_name" class="block text-sm">{user.name}</span>
					<span data-testid="account_email" class="block truncate text-sm font-medium"
						>{user.email}</span
					>
				</DropdownHeader>
				<DropdownItem on:click={signout}>ログアウト</DropdownItem>
			</Dropdown>
		{/if}
	</Navbar>

	<div class="flex">
		<slot />
	</div>
</div>
