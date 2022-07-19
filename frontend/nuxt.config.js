export default {
  colorMode: {
    classSuffix: ''
  },
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'guessr',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/tailwindcss
    "@nuxtjs/color-mode",
    '@nuxtjs/tailwindcss',
    '@braid/vue-formulate/nuxt'
  ],

  // Modules: https://o.nuxtjs.dev/config-modules
  modules: [
    '@nuxtjs/color-mode',
    '@nuxtjs/axios',
    '@nuxtjs/auth-next'
  ],
  axios: {
    baseURL: 'http://localhost:8000'
  },
  auth: {
    strategies: {
      local: {
        token: {
          property: "access_token",
          global: true,
          required: true,
          type: "Bearer"
        },
        user: {
          property: false,
          autoFetch: true
        },
        endpoints: {
          login: { url: "/auth/token/", method: "post" },
          user: { url: "/auth/users/me/", method: "get" },
          logout: false
        }
      }
    },
    redirect: {
      login: "/login",
      callback: '/',
      home: '/',
      logout: "/"
    }
  },
  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  }
}
