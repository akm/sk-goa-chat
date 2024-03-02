<script lang="ts">
	import { isFirebaseError } from '$lib/firebase';
	import { auth } from '$lib/firebase/auth';
	import { POST } from '$lib/openapi_client';
	import { page } from '$app/stores';

	import {
		type UserCredential,
		createUserWithEmailAndPassword,
		updateProfile
	} from 'firebase/auth';

	import { onMount } from 'svelte';
	import { Label, Input, Button, Heading, P } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';

	onMount(() => {
		// analytics の初期化は onMount の外で実行すると  window is not defined というエラーが firebase の内部
		// const analytics = getAnalytics(app);
	});

	let email = '';
	let password = '';
	let accountName = '';
	let errorMessage = '';

	const signup = async () => {
		let userCredential: UserCredential;

		let serviceWorkerUID: string | undefined;
		const swMessageHandler = (event: MessageEvent) => {	
			console.log("swMessageHandler event:", event.data);
			serviceWorkerUID = event.data.uid;
		};

		navigator.serviceWorker.addEventListener("message", swMessageHandler);
		try {
			try {
				userCredential = await createUserWithEmailAndPassword(auth, email, password);
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
			navigator.serviceWorker.removeEventListener("message", swMessageHandler);
		}

		const result = await POST('/api/users', { body: { email: email, name: accountName } });
		if (result.error) {
			errorMessage = result.error.message;
			return;
		}
		console.log('signup result.data', result.data);

		if (accountName != '') {
			try {
				await updateProfile(userCredential.user, { displayName: accountName });
			} catch (err) {
				console.error(err);
				// 例外をthrowしてしまうと、再度SignUpするためにはユーザーを削除するか、ユーザーが
				// 登録済みかどうかの判断が必要になります。基本的に displayName は後でユーザー
				// が変更可能なものなので、ここで設定に失敗した場合には 「未設定」 と表示すること
				// にして、ここではエラーにしません。
				// throw err;
			}
		}

		const idToken = await userCredential.user.getIdToken();
		localStorage.setItem('idToken', idToken);
		localStorage.setItem('refreshToken', userCredential.user.refreshToken);

		// await goto('/', { replaceState: true });
		window.location.href = $page.url.origin + '/';
	};

	const signupOnEnter = (e: KeyboardEvent) => {
		if (e.key === 'Enter') {
			signup();
		}
	};
	const clearErrorMessage = () => {
		errorMessage = '';
	};
</script>

<div class="space-y-4 mt-8">
	<Heading tag="h3" class="mb-4">アカウント登録</Heading>

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
			on:keypress={signupOnEnter}
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
			on:keypress={signupOnEnter}
			on:change={clearErrorMessage}
		/>
	</Label>

	<Label class="block">
		<span class="text-gray-700">アカウント名</span>
		<Input
			class="mt-1 block w-full"
			bind:value={accountName}
			placeholder="アカウント名"
			on:keypress={signupOnEnter}
			on:change={clearErrorMessage}
		/>

		<P size="xs">アカウント名は後で変更することが可能です</P>
	</Label>

	<div class="flex">
		<Button class="flex-none mt-4 h-12" on:click={signup}>アカウント登録</Button>
		<div>
			<Button class="mt-4 ml-8 h-12" color="alternative" href="/signin">ログイン</Button>
			<div class="ml-8">すでにアカウントを登録済みの場合はこちら</div>
		</div>
	</div>
</div>
