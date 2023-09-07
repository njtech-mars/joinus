<script lang="ts">
	import { browser } from '$app/environment';

	export let showModal: boolean;

	$: if (browser) {
		if (showModal) {
			const x = window.scrollX;
			const y = window.scrollY;
			window.onscroll = () => window.scrollTo(x, y);
		} else window.onscroll = () => {};
	}
</script>

<div class="backdrop" class:active={showModal}>
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<div class="modal" class:active={showModal} on:click|self={() => (showModal = false)}>
		<slot />
	</div>
</div>

<style lang="postcss">
	.backdrop {
		@apply bg-black/30 fixed hidden inset-0;
		&.active {
			@apply block;
		}
	}
	.modal {
		@apply w-full h-full flex flex-col items-center justify-center;
		&.active {
			animation: zoom 500ms cubic-bezier(0.34, 1.56, 0.64, 1);
		}
	}
	@keyframes zoom {
		from {
			transform: scale(0.9);
		}
		to {
			transform: scale(1);
		}
	}
</style>
