<script lang="ts">
	import { onMount } from 'svelte';

	import applicant from '$stores/applicant';
	import type { Status } from '$types/status';

	import Form from './components/Form/Form.svelte';
	import Message from './components/Message/Message.svelte';

	let status: Status;

	onMount(async () => {
		if (new Date() > new Date('2023-09-21T15:59:59.000Z')) status = 'expired';
		else status = localStorage.getItem('student_id') ? 'submitted' : 'idle';
	});
</script>

<svelte:head>
	<title>欢迎加入Mars工作室 • 参与报名</title>
</svelte:head>

<main>
	{#if status === 'idle'}
		<Form bind:status />
	{:else if status === 'expired'}
		<Message
			title="报名已截至"
			message="对不起，报名已截至，如果你还想加入我们可以发送邮件至 njtech-mars@outlook.com"
		/>
	{:else if status === 'done'}
		<Message
			bind:status
			title={$applicant ? '更新成功' : '提交成功'}
			message={$applicant
				? '恭喜你信息更新成功，但是请勿频繁更新信息，这可能会导致我们后台数据不同步'
				: '恭喜你提交成功，稍后会有邮件发送至你的邮箱，你也可以前往首页加群关注最新消息'}
		/>
	{:else if status === 'submitted'}
		<Message
			bind:status
			title="重复提交"
			message="你已经提交过啦，如果你还没有收到邮件通知可以前往首页加群，了解最新消息"
		/>
	{/if}
</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply flex flex-col items-center justify-center px-4 py-5 md:py-10;
	}
</style>
