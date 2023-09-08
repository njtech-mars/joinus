<script lang="ts">
	import IoIosArrowDown from 'svelte-icons/io/IoIosArrowDown.svelte';

	export let department: { name: string; detail: string };

	let collapse = true;
</script>

<div class="wrapper">
	<button type="button" on:click={() => (collapse = !collapse)}>
		<div class:active={collapse}><IoIosArrowDown /></div>
		<div>{department.name}</div>
	</button>
	<p class:hidden={collapse}>{department.detail}</p>
</div>

<style lang="postcss">
	button {
		@apply flex flex-row items-center gap-1 w-full text-left py-3 px-5;
		& div:first-of-type {
			@apply w-4 h-4 rotate-0 duration-300;
			&.active {
				@apply -rotate-90;
			}
		}
	}
	p {
		@apply bg-white border-t border-t-gray-400 px-5 py-3 whitespace-pre-wrap;
	}
	p:last-child {
		@apply rounded-b-md;
	}
	.wrapper {
		@apply border border-gray-400 bg-gray-100;
	}
	.wrapper:not(:last-of-type) {
		@apply border-b-0;
	}
	.wrapper:first-of-type {
		@apply rounded-t-md;
	}
	.wrapper:last-of-type {
		@apply rounded-b-md;
	}
</style>
