<script lang="ts">
	import Fireworks from '$components/Icons/Fireworks/Fireworks.svelte';
</script>

<div class="wrapper">
	<div class="w-36"><Fireworks /></div>
	<h1 class="text-3xl font-semibold mt-2 mb-3">欢迎加入Mars工作室</h1>
	<p class="text-gray-700 max-w-md">
		欢迎2023级的新生来到我们南京工业大学，如果你想加入我们Mars工作室，一同参与校园互联网项目的开发，又或者参与南工在线视频资源的上传，你可以填写以下表单加入我们
	</p>
	<p class="text-gray-500 mt-1 mb-3">报名时间：2023-09-04至2023-09-22</p>
	<div class="flex flex-row gap-5">
		<a href="/contact">加群了解</a>
		<a href="/form">参与报名</a>
	</div>
</div>

<style lang="postcss">
	.wrapper {
		min-height: calc(100vh - 100px);
		min-height: calc(100svh - 100px);
		@apply px-4 w-full flex flex-col items-center justify-center text-center;
	}
	a {
		@apply font-semibold py-2.5 px-7 rounded-full duration-300;
	}
	a:first-of-type {
		@apply border-2 border-blue-700 text-gray-800;
		&:hover {
			@apply text-blue-700;
		}
	}
	a:last-of-type {
		@apply text-white bg-blue-700;
		&:hover {
			@apply bg-blue-800;
		}
	}
</style>
