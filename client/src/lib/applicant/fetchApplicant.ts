import z from 'zod';

import url from '$lib/utils/url';
import { Applicant } from '$types/applicant';

export default async function fetchApplicant(student_id: string) {
	try {
		const path = `/api/applicant?student_id=${student_id}`;
		const res = await fetch(url(path), { credentials: 'include' }).then((res) => res.json());
		return z.object({ data: Applicant }).parse(res).data;
	} catch (err) {
		console.error(err);
		return null;
	}
}
