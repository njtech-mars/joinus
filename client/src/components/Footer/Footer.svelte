<script lang="ts">
	import IoIosShareAlt from 'svelte-icons/io/IoIosShareAlt.svelte';

	import { links } from './config';
</script>

<footer>
	<ul class="w-full grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-5">
		{#each links as category (category.group)}
			<li class="space-y-2">
				<h1 class="font-semibold">{category.group}</h1>
				<ul>
					{#each category.items as item (item.name)}
						<li>
							<a href={item.path} target="_blank">
								<div>{item.name}</div>
								<div class="icon"><IoIosShareAlt /></div>
							</a>
						</li>
					{/each}
				</ul>
			</li>
		{/each}
	</ul>

	<p class="copyright">
		<span>Copyright &copy; {new Date().getFullYear()} Mars工作室.</span>
		<a href="https://beian.miit.gov.cn" target="_blank">苏ICP备2022032826号</a>
	</p>
</footer>

<style lang="postcss">
	footer {
		@apply bg-gray-200 text-gray-700 px-4 md:px-32 py-10 flex flex-col gap-4;
	}
	a {
		@apply flex flex-row items-center gap-0.5 w-fit duration-300;
		&:hover {
			@apply text-blue-600;
		}
		& .icon {
			@apply block w-4;
		}
	}
	.copyright {
		@apply flex flex-col md:flex-row md:gap-2;
	}
</style>
