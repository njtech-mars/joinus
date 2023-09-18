<script lang="ts">
	import type { Status } from '$types/status';

	import Flower from '$components/Icons/Flower/Flower.svelte';

	export let title: string;
	export let message: string;
	export let status: Status | undefined = undefined;
</script>

<div>
	<Flower />
	<h1>{title}</h1>
	<p>{message}</p>
	<a href="/">返回</a>
	{#if status}
		<button type="button" on:click={() => (status = 'idle')}>表单有误，我要重新提交</button>
	{/if}
</div>

<style lang="postcss">
	div {
		@apply w-full flex flex-col items-center justify-center gap-2.5;
	}
	h1 {
		@apply text-2xl font-semibold text-gray-800;
	}
	p {
		@apply max-w-xs text-center text-gray-600;
	}
	a {
		@apply bg-blue-600 text-white py-2 px-6 font-semibold rounded-full duration-300;
		&:hover {
			@apply bg-blue-700;
		}
	}
	button {
		@apply text-xs text-gray-500;
		&:hover {
			@apply underline;
		}
	}
</style>
