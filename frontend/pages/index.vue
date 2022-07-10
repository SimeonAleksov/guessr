<template>
  <div class="flex flex-col justify-center items-center">
    <h1 class="font-mono text-4xl dark:text-secondary">
      How well do you know your <span class="text-primary">music</span>?
    </h1>
    <h2 class="font-mono text-md dark:text-secondary">
      Create an account and start listening and guessing!
    </h2>
  </div>
</template>

<script>
export default {
  name: 'Index',
  mounted() {
    console.log(this.$colorMode.preference)
  },
};
</script>
