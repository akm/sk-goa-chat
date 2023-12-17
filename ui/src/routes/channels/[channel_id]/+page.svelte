<script lang="ts">
	import { Heading, TextPlaceholder } from 'flowbite-svelte';

	import type { Channel } from '$lib/models/channel';

	export let data = {
		channel: Channel
	};
</script>

<Heading tag="h3" class="mb-4">{data.channel.name}</Heading>

<TextPlaceholder class="mb-8" />

<TextPlaceholder class="mb-8" />
