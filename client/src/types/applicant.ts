import z from 'zod';

const SubmitApplicant = z.object({
	name: z.string().min(2).max(20),
	gender: z.enum(['男', '女']),
	email: z.string().email().max(320),
	qq: z.string().min(5).max(11),
	student_id: z.string().length(12),
	college: z.string().max(50),
	major: z.string().max(50),
	introduction: z.string().min(4).max(500)
});

const Applicant = z.object({ id: z.number(), submitted_at: z.string() }).merge(SubmitApplicant);

type ApplicantType = z.TypeOf<typeof Applicant>;

export { Applicant, SubmitApplicant };
export type { ApplicantType };
