<template>

  <!--  Header / Navbar  -->
  <div v-if="getUserProfile.id !== 0">
    <Header />
  </div>
  <div v-if="getUserProfile.id === 0" class="d-flex">
    <div class="ms-auto p-3">
      <ThemeButton />
    </div>
  </div>
  <div class="container-fluid">
    <div class="row">

      <!--  Sidebar Navigation / Logged in only  -->
      <div v-if="getUserProfile.id !== 0">
        <SidebarNav></SidebarNav>
      </div>

      <!--  App View  -->
      <div :class="getUserProfile.id !== 0 ? 'col-md-9 ms-sm-auto col-lg-10 px-md-4' : ''">
        <router-view />
      </div>

    </div>
  </div>

  <!--  Footer  -->
  <div v-if="getUserProfile.id !== 0">
    <Footer />
  </div>

</template>

<script>
import Header from "./components/Header.vue"
import SidebarNav from "./components/SidebarNav";
import Footer from "./components/Footer";
import {mapGetters} from "vuex";
import ThemeButton from "./components/ThemeButton";
export default {
  name: 'App',
  components: {
    ThemeButton,
    Header,
    SidebarNav,
    Footer
  },
  computed: {
    ...mapGetters("auth", {
      getUserProfile: "getUserProfile",
    }),

  },
  mounted () {
    /*document.body.classList.add('d-flex', 'flex-column', 'h-100')*/
    document.body.classList.add('h-100')
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

a, p, i, .nav-link, label, h1, h2, h3, h4, button {
  color: var(--text-primary-color) !important;
}


body {
  font-size: .875rem;
}

.feather {
  width: 16px;
  height: 16px;
  vertical-align: text-bottom;
}



/*
 * Sidebar
 */

.sidebar {
  position: fixed;
  top: 0;
  /* rtl:raw:
  right: 0;
  */
  bottom: 0;
  /* rtl:remove */
  left: 0;
  z-index: 100; /* Behind the navbar */
  padding: 48px 0 0; /* Height of navbar */

}

@media (max-width: 767.98px) {
  .sidebar {
    top: 5rem;
  }
}

.sidebar-sticky {
  position: relative;
  top: 0;
  height: calc(100vh - 48px);
  padding-top: .5rem;
  overflow-x: hidden;
  overflow-y: auto; /* Scrollable contents if viewport is shorter than content. */
}

.sidebar .nav-link {
  font-weight: 500;
  color: #333;
}

.sidebar .nav-link .feather {
  margin-right: 4px;
  color: #727272;
}

.sidebar .nav-link.active {
  color: #2470dc;
}

.sidebar .nav-link:hover .feather,
.sidebar .nav-link.active .feather {
  color: inherit;
}

.sidebar-heading {
  font-size: .75rem;
  text-transform: uppercase;
}

/*
 * Navbar
 */

.navbar-brand {
  padding-top: .75rem;
  padding-bottom: .75rem;
  font-size: 1rem;
  color: var(--text-primary-color) !important;
  background-color: var(--background-color-primary) !important;
}

.navbar .navbar-toggler {
  top: .25rem;
  right: 1rem;
}

.navbar .form-control {
  padding: .75rem 1rem;
  border-width: 0;
  border-radius: 0;
}

.form-control {
  color: var(--text-primary-color) !important;
  background-color: var(--background-color-secondary) !important;
  border-color: var(--theme-switch-border-color) !important;
}
.form-control:hover {
  border-color: var(--theme-switch-border-color-hover) !important;
}
.dropdown-item:hover {
  background-color: var(--background-color-secondary) !important;
}


.card {
  color: var(--text-primary-color) !important;
  background-color: var(--background-color-primary) !important;
  border-color: var(--theme-switch-border-color) !important;
}
.card-header {
  color: var(--text-primary-color) !important;
  background-color: var(--background-color-secondary) !important;
  border-color: var(--theme-switch-border-color) !important;
}
.list-group-item {
  color: var(--text-primary-color) !important;
  background-color: var(--background-color-primary) !important;
  border-color: var(--theme-switch-border-color) !important;
}

</style>
