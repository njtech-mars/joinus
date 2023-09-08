import z from 'zod';

const Options = z.enum(['开发部', '运营部', '设计部', '活动部']);

const SubmitApplicant = z.object({
	name: z.string().min(2).max(20),
	gender: z.enum(['男', '女']),
	email: z.string().email().max(320),
	qq: z.string().min(5).max(11),
	student_id: z.string().length(12),
	college: z.string().max(50),
	major: z.string().max(50),
	introduction: z.string().min(4).max(500),
	first_choice: Options,
	second_choice: Options
});

const Applicant = z.object({ id: z.number(), submitted_at: z.string() }).merge(SubmitApplicant);

type ApplicantType = z.TypeOf<typeof Applicant>;

export { Applicant, SubmitApplicant };
export type { ApplicantType };
