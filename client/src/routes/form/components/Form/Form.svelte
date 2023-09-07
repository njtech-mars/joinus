<script lang="ts">
	import inputs from '$stores/inputs';
	import toasts from '$stores/toasts';
	import applicant from '$stores/applicant';
	import submitForm from './lib/submitForm';
	import scrollToFirstError from './lib/scrollToFirstError';

	import type { Status } from '$types/status';

	import Header from './components/Header/Header.svelte';
	import SubmitButton from './components/SubmitButton/SubmitButton.svelte';

	import NameInput from './components/NameInput/NameInput.svelte';
	import GenderSelect from './components/GenderSelect/GenderSelect.svelte';
	import EmailInput from './components/EmailInput/EmailInput.svelte';
	import QQInput from './components/QQInput/QQInput.svelte';
	import StudentIdInput from './components/StudentIDInput/StudentIDInput.svelte';
	import CollegeInput from './components/CollegeInput/CollegeInput.svelte';
	import MajorInput from './components/MajorInput/MajorInput.svelte';
	import IntroductionInput from './components/IntroductionInput/IntroductionInput.svelte';
	import sendEmail from './lib/sendEmail';

	export let status: Status;

	let pending = false;

	async function handleSubmit(event: SubmitEvent) {
		if (inputs.validateAll()) scrollToFirstError();
		else {
			pending = true;
			const formData = new FormData(event.target as HTMLFormElement);
			const studentId = formData.get('student_id')?.toString() || '';
			const submitResult = await submitForm(formData);
			if (submitResult.success) {
				status = 'done';
				window.scrollTo({ top: 0 });
				localStorage.setItem('student_id', studentId);

				// send email
				const sendEmailResult = await sendEmail(studentId);
				if (!sendEmailResult.success) toasts.add(sendEmailResult.message, 'failed');
			} else toasts.add(submitResult.message, 'failed');

			pending = false;
		}
	}
</script>

<form on:submit|preventDefault={handleSubmit} class="w-full max-w-2xl flex flex-col items-end gap-7">
	<Header />

	<NameInput defaultValue={$applicant?.name || ''} />
	<GenderSelect defaultValue={$applicant?.gender || '男'} />
	<EmailInput defaultValue={$applicant?.email || ''} />
	<QQInput defaultValue={$applicant?.qq || ''} />
	<StudentIdInput defaultValue={$applicant?.student_id || ''} />
	<CollegeInput defaultValue={$applicant?.college || ''} />
	<MajorInput defaultValue={$applicant?.major || ''} />
	<IntroductionInput defaultValue={$applicant?.introduction || ''} />

	<SubmitButton {pending} />
</form>
