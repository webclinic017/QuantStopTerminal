<template>
  <header>
    <nav class="navbar navbar-expand navbar-dark fixed-top bg-dark">
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
  </header>
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
        this.$router.push("/");
      }
    }
  },
}
</script>

<style scoped>

</style>