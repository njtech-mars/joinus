<script lang="ts">
	import toasts from '$stores/toasts';
	import sendEmail from '$lib/email/sendEmail';
	import fetchApplicants from '$lib/applicant/fetchApplicants';
	import type { ApplicantType } from '$types/applicant';

	import Modal from '$components/Modal/Modal.svelte';

	export let applicant: ApplicantType;

	let showModal = false;

	async function handleSendEmail() {
		showModal = false;
		const id = toasts.add('邮件发送中...', 'pending');
		const result = await sendEmail(applicant.student_id, false);
		toasts.update(id, { message: result.message, status: result.success ? 'success' : 'failed' });
		if (result.success) fetchApplicants();
	}
</script>

<Modal bind:showModal>
	<div class="dialog">
		<h1 class="text-lg font-semibold">确认发送</h1>
		{#if applicant.email_status === 'success'}
			<p class="text-center text-gray-700 whitespace-pre-wrap">该同学已经成功发送过邮件了，是否确认再次发送</p>
		{:else}
			<p>是否确认向该同学发送邮件通知</p>
		{/if}
		<div class="w-full flex flex-row gap-4">
			<button type="button" on:click={() => (showModal = false)}>取消</button>
			<button type="button" on:click={handleSendEmail}>确认</button>
		</div>
	</div>
</Modal>

<button type="button" on:click={() => (showModal = true)}>
	{#if applicant.email_status === 'none'}
		<span>还未发送</span>
	{:else if applicant.email_status === 'success'}
		<span class="text-green-600">发送成功</span>
	{:else}
		<span class="text-red-600">发送失败</span>
	{/if}
</button>

<style lang="postcss">
	.dialog {
		@apply bg-white rounded-xl w-full max-w-xs p-7 flex flex-col items-center gap-4;
	}
	.dialog button {
		@apply w-full py-2 rounded-md;
	}
	.dialog button:first-of-type {
		@apply border border-gray-400 text-gray-700 duration-300;
		&:hover {
			@apply border-black text-black;
		}
	}
	.dialog button:last-of-type {
		@apply bg-blue-600 text-white duration-300;
		&:hover {
			@apply bg-blue-700;
		}
	}
</style>
