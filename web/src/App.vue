<template>

  <!--  Header / Navbar  -->
  <Header />

  <!--  Main  -->
  <main class="d-flex align-content-stretch">
    <router-view />
  </main>

  <!--  Footer  -->
  <Footer />
</template>

<script>
import Header from "./components/Header.vue"
import Footer from "./components/Footer";
export default {
  name: 'App',
  components: {
    Header,
    Footer
  },

  mounted () {
    document.body.classList.add('d-flex', 'flex-column', 'h-100')
    document.documentElement.classList.add('h-100')
  },
}
</script>

<style>
body {
  background-color: var(--background-color-primary) !important;
}

main {
  background-color: var(--background-color-primary);
  margin-top: 55px;
  margin-bottom: auto;
}
footer {
  background-color: var(--background-color-primary);
}

/* Define styles for the default root window element */
:root {
  --background-color-primary: #ffffff;
  --background-color-secondary: #fafafa;
  --accent-color: #cacaca;
  --text-primary-color: #222;
  --element-size: 4rem;
  --theme-switch-background-color: #f1f1f1; /* Theme switch background color */
  --theme-switch-border-color: rgb(60 60 60 / 29%);
  --theme-switch-border-color-hover: rgba(59, 59, 59, 0.29);
}

/* Define styles for the root window with dark - mode preference */
:root.dark-theme {
  --background-color-primary: #1a1a1a;
  --background-color-secondary: #2d2d30;
  --accent-color: #3f3f3f;
  --text-primary-color: rgb(255 255 255 / 87%);
  --theme-switch-background-color: #2f2f2f; /* Theme switch background color */
  --theme-switch-border-color: rgb(84 84 84 / 65%);
  --theme-switch-border-color-hover: rgba(196, 193, 193, 0.29);
}

a, p, .nav-link {
  color: var(--text-primary-color) !important;
}

</style>
