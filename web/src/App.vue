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
  --background-color-primary: #fff;
  --background-color-secondary: #fafafa;
  --accent-color: #cacaca;
  --text-primary-color: #222;
  --element-size: 4rem;
}

/* Define styles for the root window with dark - mode preference */
:root.dark-theme {
  --background-color-primary: #1e1e1e;
  --background-color-secondary: #2d2d30;
  --accent-color: #3f3f3f;
  --text-primary-color: #ddd;
}

p {
  color: var(--text-primary-color);
}

.container {
  background-color: var(--background-color-primary);
  /*height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;*/
}

</style>
