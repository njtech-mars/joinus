<script lang="ts">
	import { onMount } from 'svelte';
	import applicant from '$stores/applicant';
	import fetchApplicant from '$lib/applicant/fetchApplicant';

	import Spinner from '$components/Spinner/Spinner.svelte';

	export let pending: boolean;
	export let confirm: boolean;

	onMount(async () => {
		const student_id = localStorage.getItem('student_id');
		if (!student_id) return;
		const data = await fetchApplicant(student_id);
		if (data) applicant.set(data);
	});
</script>

<div class="flex flex-row gap-4">
	<a href="/" type="button"> 返回 </a>

	<button disabled={pending} type="button" on:click={() => (confirm = true)}>
		{#if pending}
			<Spinner />
		{/if}
		<span>提交</span>
	</button>
</div>

<style lang="postcss">
	:is(a, button) {
		@apply py-2 px-6 rounded-md duration-300;
	}
	a {
		@apply border border-gray-400 text-gray-700;
		&:hover {
			@apply border-black text-black;
		}
	}
	button {
		@apply flex flex-row items-center gap-1 bg-blue-600 text-white;
		&:hover {
			@apply bg-blue-700;
		}
		&:disabled {
			@apply bg-gray-200 cursor-not-allowed text-gray-400;
		}
	}
</style>
