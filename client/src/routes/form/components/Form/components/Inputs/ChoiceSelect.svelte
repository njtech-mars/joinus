<script lang="ts">
	import '../form.css';
	import inputs from '$stores/inputs';

	export let defaultFirstChoice: string;
	export let defaultSecondChoice: string;

	let error = '';
	let first_choice = defaultFirstChoice;
	let second_choice = defaultSecondChoice;

	$: inputs.register('first_choice', first_choice + ',' + second_choice, validate);
	$: error = $inputs.find((input) => input.id === 'first_choice')?.error || '';

	function handleChange() {
		inputs.update('first_choice', first_choice + ',' + second_choice);
	}

	function validate(value: string) {
		const [first, second] = value.split(',');
		if (first && second && first === second) return '志愿重复，请修改第一或第二志愿';
		else return '';
	}
</script>

<div class="form" class:error>
	<h1 class="title">
		<span>第一志愿</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="first_choice" class="label">工作室你最想加入的部门，我们会优先考虑你的第一志愿</label>
	<select id="first_choice" name="first_choice" class="input" bind:value={first_choice} on:change={handleChange}>
		<option value="项目部">项目部</option>
		<option value="运维部">运维部</option>
		<option value="资源部">资源部</option>
		<option value="设计部">设计部</option>
	</select>
	<p class="err-msg">{error}</p>
</div>

<div class="form" class:error>
	<h1 class="title">
		<span>第二志愿</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="second_choice" class="label">工作室你还想加入的部门，请勿与第一志愿相同</label>
	<select id="second_choice" name="second_choice" class="input" bind:value={second_choice} on:change={handleChange}>
		<option value="运维部">运维部</option>
		<option value="项目部">项目部</option>
		<option value="资源部">资源部</option>
		<option value="设计部">设计部</option>
	</select>
	<p class="err-msg">{error}</p>
</div>
