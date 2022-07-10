<template>
  <div class="flex flex-col justify-center items-center">
  <h1 class="font-mono text-4xl dark:text-secondary mb-4">
    Create an account
  </h1>
  <form class="flex flex-col justify-center items-center space-y-8">
    <Input placeholder="Username" />
    <Input placeholder="Password" type="password"/>
    <Button>
      Register
    </Button>
  </form>
  </div>
</template>

<script>
import Input from '../../components/shared/Input.vue';
import Button from '../../components/shared/Button.vue';

export default {
  name: 'register',
  components: {
    Input,
    Button,
  },
  mounted() {
    console.log(this.$colorMode.preference)
  },
};
</script>
