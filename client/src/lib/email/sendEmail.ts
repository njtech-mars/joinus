import z from 'zod';

import url from '$lib/utils/url';

export default async function sendEmail(studentId: string, once: boolean = true) {
	try {
		let path = `/api/email?student_id=${studentId}`;
		if (!once) path += '&once=false';
		const res = await fetch(url(path), { method: 'POST', credentials: 'include' });
		const message = z.object({ message: z.string() }).parse(await res.json()).message;
		return { success: res.status < 500, message };
	} catch (err) {
		console.error(err);
		return { success: false, message: '邮件发送失败，请联系我们给你重新发送' };
	}
}
