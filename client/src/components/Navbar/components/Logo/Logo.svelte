<a href="/" class="z-20 flex flex-row items-center gap-1.5">
  <div class="title">Mars工作室</div>
</a>

<style lang="postcss">
  .title {
    @apply relative px-3 py-1 isolate text-white font-semibold italic;
  }
  .title::before {
    content: "";
    border-radius: 0.8rem 0.3rem 0.8rem 0.2rem;
    @apply -z-10 absolute top-0 right-0 w-full h-full bg-blue-600;
  }
  .title::after {
    content: "";
    border-radius: 0.8rem 0.3rem 0.8rem 0.2rem;
    @apply -z-20 absolute -top-0.5 -right-0.5 w-full h-full bg-blue-400;
  }
</style>
