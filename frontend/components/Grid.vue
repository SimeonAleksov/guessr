<template>
  <div class="grid grid-cols-4 gap-4">
    <div v-for="album in albums" :key="album.ID" class="max-w-xs">
      <img class="object-cover w-28 h-28 scale-100 hover:scale-150 ease-in duration-300" :src="getImgUrl(album.ImagePath.String)" :alt="album.Name" @click="onClick" :id="album.ID" />
      <h1 class="font-mono text-md dark:text-secondary mt-2 ">
        {{ album.Name }}
      </h1>
    </div>
  </div>
</template>

<script>
export default {
  name: "Grid",
  props: {
    albums: {
      type: Array,
      required: true,
    }
  },
  methods: {
    getImgUrl(img) {
      return require('../assets/albums/' + img + '.jpg')
    },
    onClick(e) {
      this.$emit('click', e.target.id)
    }
  }
}
</script>

<style scoped>
</style>
