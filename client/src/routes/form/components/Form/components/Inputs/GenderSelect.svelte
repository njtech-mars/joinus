<script lang="ts">
	import '../form.css';

	export let defaultValue: '男' | '女';
</script>

<div class="form">
	<h1 class="title">
		<span>性别</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="gender" class="label">你的性别</label>
	<select id="gender" name="gender" class="input" value={defaultValue || '女'}>
		<option value="男">男</option>
		<option value="女">女</option>
	</select>
</div>
