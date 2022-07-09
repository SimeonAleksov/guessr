<template>
  <div>
    <h1 class="dark:text-primary text-4xl"> Guessr </h1>
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
