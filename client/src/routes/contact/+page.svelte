<script lang="ts">
	import qrcode from './qrcode.jpg';
</script>

<svelte:head>
	<title>欢迎加入Mars工作室 • 加群了解</title>
</svelte:head>

<main>
	<img src={qrcode} alt="qrcode" />
</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply flex flex-col items-center justify-center px-4;
	}
	img {
		box-shadow: 0 0 20px #d1d5db;
		@apply rounded-xl w-full max-w-[270px] md:max-w-xs;
	}
</style>
