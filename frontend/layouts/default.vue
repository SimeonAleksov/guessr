<template>
    <div class="dark:bg-onyx dark:text-onyx h-screen flex flex-col">
      <Navbar />
      <nuxt class="mx-auto p-4 flex-grow" />
      <Footer />
    </div>
</template>

<script>
import Button from '../components/shared/Button.vue';
import Navbar from '../components/Navbar.vue';
import Footer from '../components/Footer.vue';

export default {
    components: {
      Button,
      Navbar,
      Footer,
    }
};
</script>
