<script lang="ts">
	import { isFirebaseError } from '$lib/firebase';
	import { auth } from '$lib/firebase/auth';
	import { page } from '$app/stores';

	import { onMount } from 'svelte';
	import { signInWithEmailAndPassword, type UserCredential } from 'firebase/auth';

	import { Label, Input, Button, Heading } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';

	onMount(() => {
		// const analytics = getAnalytics(app);
	});
	let email = '';
	let password = '';
	let errorMessage = '';

	const signin = async () => {
		let userCredential: UserCredential;

		let serviceWorkerUID: string | undefined;
		const swMessageHandler = (event: MessageEvent) => {
			console.log('swMessageHandler event:', event.data);
			serviceWorkerUID = event.data.uid;
		};
		navigator.serviceWorker.addEventListener('message', swMessageHandler);
		try {
			try {
				userCredential = await signInWithEmailAndPassword(auth, email, password);
				console.log('userCredential', userCredential);
			} catch (err) {
				if (isFirebaseError(err)) {
					errorMessage = `[${err.code}] ${err.message}`;
					return;
				}
				throw err;
			}

			for (let i = 0; i < 30; i++) {
				await new Promise((resolve) => setTimeout(resolve, 100));
				if (serviceWorkerUID === userCredential.user.uid) {
					break;
				}
			}
			if (serviceWorkerUID !== userCredential.user.uid) {
				errorMessage = 'failed to get uid from service-worker';
				return;
			}
		} finally {
			navigator.serviceWorker.removeEventListener('message', swMessageHandler);
		}

		let idToken: string;
		try {
			idToken = await userCredential.user.getIdToken();
			console.log('after getIdToken: idToken', idToken);
			localStorage.setItem('idToken', idToken);
			localStorage.setItem('refreshToken', userCredential.user.refreshToken);
		} catch (err) {
			console.error(err);
			errorMessage = 'failed to get token';
			throw err;
		}

		// await goto('/', { replaceState: true });
		window.location.href = $page.url.origin + '/';
	};

	const signinOnEnter = (e: KeyboardEvent) => {
		if (e.key === 'Enter') {
			signin();
		}
	};
	const clearErrorMessage = () => {
		errorMessage = '';
	};
</script>

<div class="space-y-4 mt-8">
	<Heading tag="h3" class="mb-4">ログイン</Heading>

	{#if errorMessage}
		<div class="flex">
			<ExclamationCircleOutline class="text-red-500 h-6 w-6 flex-none" />
			<p class="text-red-500 ml-2">
				{errorMessage}
			</p>
		</div>
	{/if}

	<Label class="block">
		<span class="text-gray-700">メールアドレス</span>
		<Input
			class="mt-1 block w-full"
			bind:value={email}
			placeholder="メールアドレス"
			on:keypress={signinOnEnter}
			on:change={clearErrorMessage}
		/>
	</Label>
	<Label class="block">
		<span class="text-gray-700">パスワード</span>
		<Input
			class="mt-1 block w-full"
			type="password"
			bind:value={password}
			placeholder="パスワード"
			on:keypress={signinOnEnter}
			on:change={clearErrorMessage}
		/>
	</Label>
	<div class="flex">
		<Button class="flex-none mt-4 h-12" on:click={signin}>ログイン</Button>
		<div>
			<Button class="mt-4 ml-8 h-12" color="alternative" href="/signup">アカウント登録</Button>
			<div class="ml-8">まだアカウント登録をされていない場合はこちら</div>
		</div>
	</div>
</div>
