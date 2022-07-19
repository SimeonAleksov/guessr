<template>
  <input
    :placeholder="placeholder"
    :type="type"
    :model="model"
    @input="onInput"
    @change="onChange"
  />
</template>
<script>
export default {
  name: 'Input',
  props: {
    type: {
      type: String,
      default: () => 'text',
    },
    placeholder: {
      type: String,
      required: true,
    },
    model: {
      required: true,
    }
  },
  methods: {
    onInput(e) {
      this.$emit("input", e.target.value);
    },
    onChange(e) {
      this.$emit("change", e.target.value);
    },
  }
}
</script>
<style lang="scss">
label {
  color: #d3d0cb;
  margin-top: 1rem;
}
input {
  width: 200px;
  display: block;
  border: none;
  padding: 10px 0;
  border-bottom: solid 1px #1abc9c;
  -webkit-transition: all 0.3s cubic-bezier(0.64, 0.09, 0.08, 1);
  transition: all 0.3s cubic-bezier(0.64, 0.09, 0.08, 1);
  background: -webkit-linear-gradient(top, rgba(255, 255, 255, 0) 96%, #e2c044 4%);
  background: linear-gradient(to bottom, rgba(255, 255, 255, 0) 96%, #e2c044 4%);
  background-position: -200px 0;
  background-size: 200px 100%;
  background-repeat: no-repeat;
  color: #e2c044;
}

input:focus, input:valid {
 box-shadow: none;
 outline: none;
 background-position: 0 0;
}


input::-webkit-input-placeholder {
 -webkit-transition: all 0.3s ease-in-out;
 transition: all 0.3s ease-in-out;
}

input:focus::-webkit-input-placeholder, input:valid::-webkit-input-placeholder {
 color: #e2c044;
 font-size: 11px;
 -webkit-transform: translateY(-8px);
 transform: translateY(-8px);
 visibility: visible !important;
}
</style>
