<script lang="ts">
	import { Heading, TextPlaceholder } from 'flowbite-svelte';

	import type { Channel } from '$lib/models/channel';
	import { Button, Modal, Label, Input, Hr } from 'flowbite-svelte';
</script>

<Heading tag="h3" class="mb-4">New Channel</Heading>

<form class="flex flex-col space-y-6" action="#">
	<Label class="space-y-2">
		<span>Name</span>
		<Input type="text" name="name" placeholder="new channel name" required />
	</Label>
	<div class="flex">
		<Button type="submit" class="mr-4">Create</Button>
	</div>
</form>
