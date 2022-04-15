<template>

  <header class="qst-header-nav navbar sticky-top flex-md-nowrap p-0">
    <router-link to="/" class="navbar-brand col-md-3 col-lg-2 me-0 px-3">QuantstopTerminal</router-link>
    <button class="navbar-toggler position-absolute d-md-none" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
      <svg class="m-auto" id="menu-bars" viewBox="0 0 20 20" width="20" height="20" fill="currentColor">
        <rect width="20" height="4"></rect>
        <rect y="6.5" width="20" height="4"></rect>
        <rect y="13" width="20" height="4"></rect>
      </svg>
    </button>
    <input v-if="getUserProfile.id !== 0" class="form-control w-100 me-3" type="text" placeholder="Search" aria-label="Search">
    <div class="d-flex qst-collapse collapse">
        <span class="nav-item m-auto me-3">
          <theme-button />
        </span>

      <div class="dropdown me-2" v-if="getUserProfile.id !== 0">
        <a href="#" class="d-flex align-items-center text-decoration-none dropdown-toggle" id="userDropdown" data-bs-toggle="dropdown" aria-expanded="false">
          <img src="//ssl.gstatic.com/accounts/ui/avatar_2x.png" alt="" width="32" height="32" class="rounded-circle me-2">
          <strong>{{ getUserProfile.username }}</strong>
        </a>
        <ul class="qst-user-dropdown dropdown-menu dropdown-menu-end text-small shadow" aria-labelledby="userDropdown">
          <li><router-link to="/settings" class="dropdown-item">Settings</router-link></li>
          <li><router-link to="/profile" class="dropdown-item">Profile</router-link></li>
          <li><hr class="dropdown-divider"></li>
          <li><a class="dropdown-item" @click.prevent="logOut" href="/logout"><font-awesome-icon icon="sign-out-alt" /> Sign out</a></li>
        </ul>
      </div>
    </div>
  </header>

<!--  <header class="navbar navbar-expand-md sticky-top flex-md-nowrap p-0 qst-header-nav">

      &lt;!&ndash; Logo &ndash;&gt;
      <router-link to="/" class="navbar-brand col-md-3 col-lg-2 me-0 px-3">QuantstopTerminal</router-link>

      <button v-if="getUserProfile.id !== 0" id="qst-menu-toggle" class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
&lt;!&ndash;        <span id="qst-menu-toggle-icon" class="navbar-toggler-icon"></span>&ndash;&gt;
        <svg class="m-auto" id="menu-bars" viewBox="0 0 20 20" width="20" height="20" fill="currentColor">
          <rect width="20" height="4"></rect>
          <rect y="6.5" width="20" height="4"></rect>
          <rect y="13" width="20" height="4"></rect>
        </svg>
      </button>

      &lt;!&ndash; Left Aligned Buttons &ndash;&gt;
      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">

        </li>
      </ul>

      &lt;!&ndash; Right Aligned Buttons / User Menu &ndash;&gt;
      <div class="d-flex">
        <span class="nav-item m-auto me-3">
          <theme-button />
        </span>

        <div class="dropdown me-2" v-if="getUserProfile.id !== 0">
          <a href="#" class="d-flex align-items-center text-decoration-none dropdown-toggle" id="userDropdown" data-bs-toggle="dropdown" aria-expanded="false">
            &lt;!&ndash; todo: user images ... &ndash;&gt;
            <img src="//ssl.gstatic.com/accounts/ui/avatar_2x.png" alt="" width="32" height="32" class="rounded-circle me-2">
            <strong>{{ getUserProfile.username }}</strong>
          </a>
          <ul class="qst-user-dropdown dropdown-menu dropdown-menu-end text-small shadow" aria-labelledby="userDropdown">
            <li><a class="dropdown-item" href="#">Settings</a></li>
            <li><a class="dropdown-item" href="#">Profile</a></li>
            <li><hr class="dropdown-divider"></li>
            <li><a class="dropdown-item" @click.prevent="logOut" href="/logout"><font-awesome-icon icon="sign-out-alt" /> Sign out</a></li>
          </ul>
        </div>
      </div>


  </header>-->
</template>

<script>
import {mapActions, mapGetters, mapMutations} from "vuex";
import ThemeButton from "./ThemeButton.vue"
export default {
  name: "Header",
  components: {
    ThemeButton,
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
      await this.userLogout();
      if (this.getLogout) {
        const resetUser = {
          id: 0,
          username: "",
          roles: null,
        };
        this.setUserProfile(resetUser);
        this.setLogout(false);
        await this.$router.push("/");
      }
    }
  },
}
</script>

<style scoped>
.qst-header-nav {
  color: var(--text-primary-color) !important;
  background-color: var(--background-color-primary) !important;
  border-bottom: 1px solid var(--theme-switch-border-color);
}
.navbar-brand {
  padding-top: 0.75rem;
  padding-bottom: 0.75rem;
  font-size: 1rem;
  text-align: center;
  /*border-right: 1px solid var(--theme-switch-border-color);*/
}
.qst-user-dropdown {
  color: var(--text-primary-color) !important;
  background-color: var(--background-color-primary) !important;
}
.navbar-toggler {
  height: 30px;
  width: 30px;
  text-align: center;
  opacity: 1;
  color: var(--text-primary-color) !important;
  background-color: var(--theme-switch-background-color) !important;
}
#menu-bars {
  text-align: center;
  position: relative;
  right: 8px;
  bottom: 2px;
}
</style>