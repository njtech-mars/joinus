<script lang="ts">
	import '../form.css';
	import inputs from '$stores/inputs';

	export let defaultValue: string;

	let error = '';

	$: inputs.register('introduction', defaultValue, validate);
	$: error = $inputs.find((input) => input.id === 'introduction')?.error || '';

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		inputs.update('introduction', value);
	}

	function validate(value: string) {
		if (value.length === 0) return '自我介绍不能为空';
		else if (value.length < 4) return '你的自我介绍太短啦';
		else if (value.length > 500) return '你的自我介绍太长啦';
		else return '';
	}
</script>

<div class="form" class:error>
	<h1 class="title">
		<span>自我介绍</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="introduction" class="label">写写你擅长的领域、个人爱好、个人主页等等，让我们对你有更好的了解</label>
	<textarea
		id="introduction"
		name="introduction"
		class="input h-32"
		placeholder="输入你的姓名"
		value={defaultValue}
		on:change={handleChange}
	/>
	<p class="err-msg">{error}</p>
</div>
