<script lang="ts">
	import { Button, Listgroup } from 'flowbite-svelte';

	import '../../app.pcss';
	import { page } from '$app/stores';
	import type { Channel } from '$lib/models/channel';

	const channelLinks = ($page.data.channels as Channel[]).map((channel) => ({
		name: channel.name,
		href: `/channels/${channel.id}`
	}));
</script>

<div class="flex flex-col" data-testid="channel_list_pane">
	{#if channelLinks.length > 0}
		<Listgroup active items={channelLinks} let:item class="w-48" data-testid="channel_list">
			{item.name}
		</Listgroup>
	{/if}

	<Button class="mt-4" color="alternative" href="/channels/new">New Channel</Button>
</div>

<div class="grow m-4">
	<slot />
</div>
