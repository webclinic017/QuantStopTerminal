<template>
  <div id="app">

    <nav class="navbar navbar-expand navbar-dark bg-dark">
      <div class="container-fluid">
        <router-link to="/" class="navbar-brand nav-link">QuantstopTerminal</router-link>

        <ul class="navbar-nav me-auto">
          <li class="nav-item">
            <router-link to="/home" class="nav-link">
              <font-awesome-icon icon="home" /> Home
            </router-link>
          </li>
          <li v-if="showAdminBoard" class="nav-item">
            <router-link to="/admin" class="nav-link">Admin Board</router-link>
          </li>
          <li v-if="showModeratorBoard" class="nav-item">
            <router-link to="/mod" class="nav-link">Moderator Board</router-link>
          </li>
          <li class="nav-item">
            <router-link v-if="getUserProfile.id !== 0" to="/user" class="nav-link">User</router-link>
          </li>
        </ul>

        <div class="nav-item">
          <theme-button />
        </div>

        <div v-if="getUserProfile.id === 0" class="navbar-nav">
          <li class="nav-item">
            <router-link to="/register" class="nav-link">
              <font-awesome-icon icon="user-plus" /> Sign Up
            </router-link>
          </li>
          <li class="nav-item">
            <router-link to="/login" class="nav-link">
              <font-awesome-icon icon="sign-in-alt" /> Login
            </router-link>
          </li>
        </div>

        <div v-if="getUserProfile.id !== 0" class="navbar-nav">
          <li class="nav-item">
            <router-link to="/profile" class="nav-link">
              <font-awesome-icon icon="user" />
              {{ getUserProfile.username }}
            </router-link>
          </li>
          <li class="nav-item">
            <a class="nav-link" @click.prevent="logOut" href="/logout">
              <font-awesome-icon icon="sign-out-alt" /> LogOut
            </a>
          </li>
        </div>

      </div>
    </nav>

    <div class="container">
      <router-view />
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from "vuex";
import ThemeButton from "./components/ThemeButton.vue"
export default {
  name: 'App',
  components: {
    ThemeButton
  },
  computed: {
    ...mapGetters("auth", {
      getUserProfile: "getUserProfile",
      getLogout: "getLogout",
    }),
    showAdminBoard() {
      if (this.getUserProfile && this.getUserProfile['roles']) {
        return this.getUserProfile['roles'].includes('admin');
      }

      return false;
    },
    showModeratorBoard() {
      if (this.getUserProfile && this.getUserProfile['roles']) {
        return this.getUserProfile['roles'].includes('moderator');
      }

      return false;
    }
  },
  methods: {
    ...mapActions("auth", {
      userLogout: "userLogout",
    }),
    ...mapMutations("auth", {
      setLogout: "setLogout",
      setUserProfile: "setUserProfile",
    }),
    async logOut() {
      /*this.$store.dispatch('auth/logout');
      this.$router.push('/login');*/
      await this.userLogout();
      if (this.getLogout) {
        const resetUser = {
          id: 0,
          username: "",
          roles: null,
          //email: "",
          //phone: "",
        };
        this.setUserProfile(resetUser);
        this.setLogout(false);
        this.$router.push("/");
      }
    }
  }
}
</script>

<style>
#app {
  background-color: var(--background-color-primary);
}
html,
body {
  background-color: var(--background-color-primary);
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

}
/* Define styles for the default root window element */
:root {
  --background-color-primary: #ebebeb;
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
