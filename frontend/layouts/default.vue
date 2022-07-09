<template>
    <div class="dark:bg-onyx dark:text-onyx h-screen flex flex-col">
      <select
        v-model="$colorMode.preference"
        class="border w-24 h-8 dark:bg-gray-900 dark:text-white dark:border-gray-700"
      >
        <option value="light">Light</option>
        <option value="dark">Dark</option>
      </select>
      <nuxt class="mx-auto p-4 flex-grow" />
    </div>
</template>

<script>
export default {
  head: {
    title: "Nuxt.js with TailwindCSS",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "description",
        content: "Nuxt.js with TailwindCSS example",
      },
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
  },
};
</script>
