<script lang="ts">
	import '../form.css';
	import inputs from '$stores/inputs';

	export let defaultValue: string;

	let error = '';

	$: inputs.register('qq', defaultValue, validate);
	$: error = $inputs.find((input) => input.id === 'qq')?.error || '';

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		inputs.update('qq', value);
	}

	function validate(value: string) {
		if (value.length === 0) return 'QQ号不能为空';
		else if (value.length < 5) return '你的QQ号太短啦';
		else if (value.length > 11) return '你的QQ号太长啦';
		else return '';
	}
</script>

<div class="form" class:error>
	<h1 class="title">
		<span>QQ</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="qq" class="label">我们联系你的备用方案，请填写正确的QQ号</label>
	<input
		id="qq"
		name="qq"
		type="number"
		class="input"
		placeholder="输入你的QQ号"
		value={defaultValue}
		on:change={handleChange}
	/>
	<p class="err-msg">{error}</p>
</div>
