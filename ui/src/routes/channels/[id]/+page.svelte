<script lang="ts">
	import { Heading, TextPlaceholder, Alert } from 'flowbite-svelte';
	import { Button, Modal, Label, Input, Hr } from 'flowbite-svelte';
	import { CogOutline, TrashBinSolid, InfoCircleSolid } from 'flowbite-svelte-icons';

	import type { Channel } from '$lib/models/channel';

	export let data = {
		channel: Channel
	};

	let name = data.channel.name;
	let errorMessage = '';

	let settingVisible = false;

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
</script>

<Heading tag="h3" class="mb-4">
	{data.channel.name}
	<Button pill={true} outline={true} class="!p-2" on:click={() => (settingVisible = true)}>
		<CogOutline class="w-4 h-4 text-primary-700" />
	</Button>
</Heading>

<TextPlaceholder class="mb-8" />

<TextPlaceholder class="mb-8" />

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
