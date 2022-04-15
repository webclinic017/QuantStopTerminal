<template>
  <table class="table table-striped table-bordered table-hover">
    <thead class="thead">
    <tr>
      <th>#</th>
      <th>Username</th>
      <th>Password</th>
      <th>Salt</th>
      <th>Roles</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="user in users">
      <td>{{user.ID}}</td>
      <td>{{user.Username}}</td>
      <td>{{user.Password}}</td>
      <td>{{user.Salt}}</td>
      <td><span v-for="role in user.Roles" class="badge bg-primary">{{ role }}</span></td>
    </tr>
    </tbody>
  </table>
</template>

<script>
import {mapActions, mapGetters} from "vuex";

export default {
  name: "UserManager",
  computed: {
    ...mapGetters("admin", {
      users: "getUsers",
    }),
  },
  methods: {
    ...mapActions("admin", {
      actionGetUsers: "getAllUsers",
    }),
    async getAll() {
      await this.actionGetUsers().then(
          (res) => {
            //this.subsystems = res
          },
          (error) => {
            this.message = error.toString() + " | " + error.response.status;
          }
      );
    },
  },
  beforeMount() {
    this.getAll();
  },
}
</script>

<style scoped>

</style>