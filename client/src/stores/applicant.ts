import { writable } from 'svelte/store';
import type { ApplicantType } from '$types/applicant';

const applicant = writable<ApplicantType | null>(null);

export default applicant;
