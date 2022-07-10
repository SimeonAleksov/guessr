<template>
  <button
    v-on="$listeners"
    class="button-64 hover:text-onyx active:text-onyx focus:text-onyx"
    role="button"
  >
      <span class="text text-center flex justify-center items-center text-primary hover:text-onyx active:text-onyx focus:text-onyx">
        <slot />
      </span>
  </button>
</template>
<script>
export default {
  name: 'TriviaButton',
}
</script>

<style>
.button-64 {
  align-items: center;
  background-image: linear-gradient(144deg,#587b7f, #e2c044 50%,#d3d0cb);
  border: 0;
  border-radius: 8px;
  box-shadow: rgba(151, 65, 252, 0.2) 0 15px 30px -5px;
  box-sizing: border-box;
  color: #FFFFFF;
  display: flex;
  font-family: Phantomsans, sans-serif;
  font-size: 20px;
  justify-content: center;
  line-height: 1em;
  max-width: 100%;
  min-width: 140px;
  padding: 3px;
  text-decoration: none;
  user-select: none;
  -webkit-user-select: none;
  touch-action: manipulation;
  white-space: nowrap;
  cursor: pointer;
}

.button-64:active,
.button-64:hover {
  outline: 0;
  color: #587b7f;
  background-image: linear-gradient(144deg,#587b7f, #e2c044 50%,#d3d0cb);
}


.button-64 span {
  background-color: #1e2019;
  padding: 16px 24px;
  border-radius: 6px;
  width: 100%;
  height: 100%;
  transition: 300ms;
}

.button-64:hover span {
  background: none;
}


.button-64:active span {
  background: none;
  color: #393e41;
}


.button-64:focus span {
  background: none;
  color: #393e41;
}

@media (min-width: 768px) {
  .button-64 {
    font-size: 24px;
    min-width: 196px;
  }
}
</style>
