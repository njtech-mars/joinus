<script lang="ts">
	import Modal from '$components/Modal/Modal.svelte';

	export let confirm: boolean;
</script>

<Modal bind:showModal={confirm}>
	<div class="dialog">
		<h1>确认提交</h1>
		<p>是否确认提交，请确保你的邮箱可用，学号的正确性</p>
		<div class="w-full flex flex-row gap-4">
			<button type="button" on:click={() => (confirm = false)}>去检查</button>
			<button type="submit" on:click={() => (confirm = false)}>我要提交</button>
		</div>
	</div>
</Modal>

<style lang="postcss">
	.dialog {
		@apply w-full max-w-xs bg-white rounded-xl p-7 flex flex-col items-center gap-4;
	}
	h1 {
		@apply font-semibold text-lg;
	}
	p {
		@apply text-center text-gray-600;
	}
	button {
		@apply w-full py-2 rounded-md;
	}
	button[type='button'] {
		@apply border border-gray-400 text-gray-700 duration-300;
		&:hover {
			@apply border-black text-black;
		}
	}
	button[type='submit'] {
		@apply bg-blue-600 text-white duration-300;
		&:hover {
			@apply bg-blue-700;
		}
	}
</style>
