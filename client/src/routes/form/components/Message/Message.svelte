<script lang="ts">
	import type { Status } from '$types/status';

	import Flower from '$components/Icons/Flower/Flower.svelte';

	export let status: Status;
	export let message: string;
	export let title: string;
</script>

<div>
	<Flower />
	<h1>{title}</h1>
	<p>{message}</p>
	<a href="/">返回</a>
	<button type="button" on:click={() => (status = 'idle')}>表单有误，我要重新提交</button>
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
		@apply text-gray-800 bg-black/10 py-1 px-3 rounded-full text-sm duration-300;
		&:hover {
			@apply bg-black/20;
		}
	}
	button {
		@apply text-xs text-gray-500;
		&:hover {
			@apply underline;
		}
	}
</style>
