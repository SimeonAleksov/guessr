const colors = require("tailwindcss/colors");
export default {
  darkMode: 'class',
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./nuxt.config.{js,ts}",
  ],
  mode: "jit",
  purge: ["./src/**/*.html"],
  // dark: "media", // or 'media' or 'class'
  theme: {
    darkSelector: '.dark',
    extend: {
      colors: {
        dark: '#1e2019',
        steel: '#587b7f',
        primary: '#e2c044',
        secondary: '#d3d0cb',
        onyx: '#393e41',
      },
    },
  },
  variants: {
    extend: {
      backgroundColor: ["dark", "dark-hover", "dark-group-hover", "dark-even", "dark-odd", "hover", "responsive"],
      borderColor: ["dark", "dark-focus", "dark-focus-within", "hover", "responsive"],
      textColor: ["dark", "dark-hover", "dark-active", "hover", "responsive"]
    },
  },
  plugins: [
  ],
};
