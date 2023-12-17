<script lang="ts">
	import { Listgroup } from 'flowbite-svelte';
	import type { Channel } from '$lib/models/channel';
	export let data: { channels: Channel[] };

	const channelLinks = data.channels.map((channel) => ({
		name: channel.name,
		href: `/channels/${channel.id}`
	}));
</script>

<div class="centered">
	<Listgroup active items={channelLinks} let:item class="w-48">
		{item.name}
	</Listgroup>
</div>
