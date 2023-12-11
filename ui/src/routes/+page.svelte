<script lang="ts">
	import { Input, Label, Checkbox, Button } from 'flowbite-svelte';
	import { TrashBinSolid } from 'flowbite-svelte-icons';
	import type { Channel } from '$lib/models/channel';
	export let data: { channels: Channel[] };
</script>

<div class="centered">
	<h1>Channels</h1>

	<ul class="todos">
		{#each data.channels as channel (channel.id)}
			<li>
				<span>{channel.name}</span>
			</li>
		{/each}
	</ul>
</div>
