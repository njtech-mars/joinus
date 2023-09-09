<script lang="ts">
	import inputs from '$stores/inputs';
	import toasts from '$stores/toasts';
	import submitForm from './lib/submitForm';
	import applicant from '$stores/applicant';
	import sendEmail from '$lib/email/sendEmail';
	import scrollToFirstError from './lib/scrollToFirstError';

	import type { Status } from '$types/status';

	import Popup from './components/Popup/Popup.svelte';
	import Header from './components/Header/Header.svelte';
	import SubmitButton from './components/SubmitButton/SubmitButton.svelte';

	import NameInput from './components/Inputs/NameInput.svelte';
	import GenderSelect from './components/Inputs/GenderSelect.svelte';
	import EmailInput from './components/Inputs/EmailInput.svelte';
	import QQInput from './components/Inputs/QQInput.svelte';
	import StudentIdInput from './components/Inputs/StudentIDInput.svelte';
	import CollegeInput from './components/Inputs/CollegeInput.svelte';
	import MajorInput from './components/Inputs/MajorInput.svelte';
	import IntroductionInput from './components/Inputs/IntroductionInput.svelte';
	import ChoiceSelect from './components/Inputs/ChoiceSelect.svelte';

	export let status: Status;

	let confirm = false;
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
				const sendEmailResult = await sendEmail(studentId, true);
				if (!sendEmailResult.success) toasts.add(sendEmailResult.message, 'failed');
			} else toasts.add(submitResult.message, 'failed');

			pending = false;
		}
	}
</script>

<form on:submit|preventDefault={handleSubmit} class="w-full max-w-2xl flex flex-col items-end gap-7">
	<Header />
	<Popup bind:confirm />

	<NameInput defaultValue={$applicant?.name || ''} />
	<GenderSelect defaultValue={$applicant?.gender || '男'} />
	<EmailInput defaultValue={$applicant?.email || ''} />
	<QQInput defaultValue={$applicant?.qq || ''} />
	<StudentIdInput defaultValue={$applicant?.student_id || ''} />
	<CollegeInput defaultValue={$applicant?.college || ''} />
	<MajorInput defaultValue={$applicant?.major || ''} />
	<ChoiceSelect
		defaultFirstChoice={$applicant?.first_choice || '开发部'}
		defaultSecondChoice={$applicant?.second_choice || '运营部'}
	/>
	<IntroductionInput defaultValue={$applicant?.introduction || ''} />

	<SubmitButton {pending} bind:confirm />
</form>
