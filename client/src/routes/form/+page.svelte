<script lang="ts">
	import { onMount } from 'svelte';

	import applicant from '$stores/applicant';
	import type { Status } from '$types/status';

	import Form from './components/Form/Form.svelte';
	import Message from './components/Message/Message.svelte';

	let status: Status;

	onMount(async () => {
		const student_id = localStorage.getItem('student_id');
		status = student_id ? 'submitted' : 'idle';
	});
</script>

<svelte:head>
	<title>欢迎加入Mars工作室 • 参与报名</title>
</svelte:head>

<main>
	{#if status === 'idle'}
		<Form bind:status />
	{:else if status === 'done'}
		<Message
			bind:status
			title={applicant ? '更新成功' : '提交成功'}
			message={applicant
				? '恭喜你信息更新成功，但是请勿频繁更新信息，这可能会导致我们后台数据不同步'
				: '恭喜你提交成功，我们会尽快通过邮件的方式通知你面试事宜，你也可以加群和我们反馈'}
		/>
	{:else if status === 'submitted'}
		<Message
			bind:status
			title="欢迎报名Mars"
			message="你的提交我们已经收到啦，如果你还没有收到邮件通知可以在群里和我们反馈"
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
