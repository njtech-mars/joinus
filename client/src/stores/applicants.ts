import { writable } from 'svelte/store';
import type { ApplicantType } from '$types/applicant';

const applicants = writable<ApplicantType[] | null>(null);

export default applicants;
