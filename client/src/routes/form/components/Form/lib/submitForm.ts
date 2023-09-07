import z from 'zod';

import url from '$lib/utils/url';

export default async function submitForm(formData: FormData) {
	try {
		const res = await fetch(url('/api/applicant'), { method: 'POST', body: formData }).then((res) => res.json());
		return z.object({ success: z.boolean(), message: z.string() }).parse(res);
	} catch (err) {
		console.error(err);
		return { success: false, message: '提交失败，请稍后重试或联系我们' };
	}
}
