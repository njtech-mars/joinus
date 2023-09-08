<script lang="ts">
	import Collapse from './Collapse.svelte';

	const departments = [
		{
			name: '开发部',
			detail: `1.有一定的前后端开发经验，代码风格良好\n2.负责项目开发和维护，工作室项目采用前后端分离架构\n3.有强烈的学习意愿和足够的时间，能够积极学习新技术`
		},
		{
			name: '运营部',
			detail:
				'1.为南工在线上传各类电影、剧集等资源\n2.了解服务器硬件组成，软件维护，会使用 Linux 操作系统。\n3.了解容器部署方案( Docker/Kubernetes )，会使用 nginx 配置服务。'
		},
		{
			name: '设计部',
			detail:
				'1.负责设计素材、美化剪辑活动照片、视频等\n2.负责推送、网站、banner、海报等平面设计工作\n3.有良好的设计审美，掌握ps、ai等软件的同学优先'
		},
		{
			name: '活动部',
			detail:
				'1.负责协助运营官方公众号“南京工业大学信息化”\n2.负责推文编辑，为大家展示自己作品，提供实践平台\n3.负责工作室团建、培训、摆摊等活动的安排'
		}
	];
</script>

<div class="wrapper">
	<h1 class="title">部门介绍</h1>
	<div class="w-full max-w-xl rounded-md">
		{#each departments as department (department.name)}
			<Collapse {department} />
		{/each}
	</div>
</div>

<style lang="postcss">
	.wrapper {
		@apply flex flex-col items-center gap-5 pb-20 px-4;
	}
	.title {
		@apply text-2xl font-semibold;
	}
</style>
