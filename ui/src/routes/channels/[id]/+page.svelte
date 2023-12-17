<script lang="ts">
	import { Heading, TextPlaceholder } from 'flowbite-svelte';

	import type { Channel } from '$lib/models/channel';
	import { Button, Modal, Label, Input, Hr } from 'flowbite-svelte';
	import { CogOutline, TrashBinSolid } from 'flowbite-svelte-icons';

	export let data = {
		channel: Channel
	};

	let settingVisible = false;
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
	<form class="flex flex-col space-y-6" action="#">
		<h3 class="text-xl font-medium text-gray-900 dark:text-white">Channel Settings</h3>
		<Label class="space-y-2">
			<span>Name</span>
			<Input type="text" name="name" placeholder={data.channel.name} required />
		</Label>
		<div class="flex">
			<Button type="submit" class="mr-4">Update</Button>
			<Button
				color="alternative"
				on:click={() => {
					settingVisible = false;
				}}>Cancel</Button
			>
		</div>
	</form>

	<Hr />

	<h3 class="text-xl font-medium text-red-700 dark:text-white">Danger Done</h3>

	<Button>
		<TrashBinSolid class="w-3.5 h-3.5 me-2" />
		Delete this channel forever
	</Button>
</Modal>
