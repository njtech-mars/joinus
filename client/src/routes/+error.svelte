<script lang="ts">
	import { page } from '$app/stores';
</script>

<svelte:head>
	<title>欢迎加入Mars工作室 • 页面错误</title>
</svelte:head>

<main>{$page.status} | {$page.error?.message || 'Error'}</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply w-full flex-1 flex items-center justify-center text-gray-700;
	}
</style>
