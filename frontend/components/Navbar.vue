<template>
<nav class="bg-dark">
  <div class="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
    <div class="relative flex items-center justify-between h-16">
      <div class="flex-1 flex items-center justify-center sm:items-stretch sm:justify-start">
        <NuxtLink to="/" class="flex-shrink-0 flex items-center">
          <h1 class="dark:text-primary text-2xl font-mono antialiased">Gue<span class="text-secondary">ss</span>r </h1>
        </NuxtLink>
      </div>
      <div class="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">
        <div class="ml-3 relative" v-if="!$auth.loggedIn">
          <ButtonLink to="/login" class="mr-4">
            Login
          </ButtonLink>
        <ButtonLink to="/register">
          Register
        </ButtonLink>
        </div>
        <div v-else>
          <ButtonLink to="/trivia/start">
            Start a game
          </ButtonLink>
        </div>
      </div>
    </div>
  </div>
</nav>
</template>
<<script>
import ButtonLink from './shared/ButtonLink.vue';

export default {
    name: "Navbar",
    components: {
      ButtonLink,
    }
}
</script>
