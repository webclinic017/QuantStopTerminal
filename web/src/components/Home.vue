<template>
  <div class="container py-4">
    <div class="p-5 mb-4 rounded-3 card">
      <div class="container-fluid py-5">
        <ul class="list-group">
          <li class="list-group-item">Database: <StatusIndicator :on="content.database"></StatusIndicator></li>
          <li class="list-group-item">Connection Monitor: <StatusIndicator :on="content.internet_monitor"></StatusIndicator></li>
          <li class="list-group-item">Timekeeper: <StatusIndicator :on="content.ntp_timekeeper"></StatusIndicator></li>
          <li class="list-group-item">Active Trader: <StatusIndicator :on="content.strategy"></StatusIndicator></li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from "vuex";
import StatusIndicator from "./StatusIndicator";
export default {
  name: "Home",
  components: {StatusIndicator},
  computed: {
    ...mapGetters("user", {
      content: "getPublicAll",
    }),
  },
  methods: {
    ...mapActions("user", {
      actionGetAll: "getPublicContent",
    }),
    async getAll() {
      this.loading = true;
      await this.actionGetAll().then(
          (res) => {
            this.content = res
          },
          (error) => {
            this.loading = false;
            this.message = error.toString() + " | " + error.response.status;
          }
      );
    },
  },
  beforeMount() {
    this.getAll();
  },
};
</script>
