<script lang="ts">
	import { Heading, Button, Label, Input, Alert } from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';

	let name = '';
	let errorMessage = '';

	const createChannel = async () => {
		const result = await fetch('/api/channels', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ name })
		});
		const json = await result.json();
		console.log('json', json);
		if (!json.id) {
			errorMessage = json.message;
			return;
		}
		// goto だと追加したチャンネルが一覧に反映されないので window.location.href でリロードする
		window.location.href = `/channels/${json.id}`;
	};
</script>

<Heading tag="h3" class="mb-4">New Channel</Heading>

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

<div>
	<Label class="space-y-2">
		<span>Name</span>
		<Input type="text" name="name" placeholder="new channel name" required bind:value={name} />
	</Label>
	<div class="flex mt-4">
		<Button type="submit" class="mr-4" on:click={createChannel}>Create</Button>
	</div>
</div>
