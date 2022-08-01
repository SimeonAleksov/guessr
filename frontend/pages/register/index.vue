<template>
  <div class="flex flex-col justify-center items-center" v-if="!$auth.loggedIn">
  <h1 class="font-mono text-4xl dark:text-secondary mb-4">
    Create an account
  </h1>
    <FormulateForm
      class="flex flex-col"
      v-model="values"
      :schema="schema"
      @submit="login"
    />
  </div>
</template>

<script>
import Button from '../../components/shared/Button.vue';

export default {
  name: 'register',
  components: {
    Button,
  },
  data () {
    return {
      values: {},
      schema: [
        {
          type: 'text',
          name: 'username',
          label: 'Username',
        },
        {
          type: 'password',
          name: 'password',
          label: 'Password',
        },
        {
          type: "submit",
          label: "Login",
          class: "flex justify-center mt-4",
        }
      ]
    }
  },

  methods: {
    async login() {
      try {
        await this.$axios.post('auth/users/', this.values, {
          headers: {
            'Content-Type': 'application/json',
          }
        })
        await this.$auth.loginWith('local', {data: this.values})
        await this.$router.push('/trivia');
      } catch (err) {
        console.log(err)
      }
    }
  }
};
</script>
<style lang="scss">
.formulate-input-error {
  color: #e2c044;
}

</style>
