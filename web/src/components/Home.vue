<template>
  <div class="container py-4">
    <div class="p-5 mb-4 rounded-3 card">
      <div class="container-fluid py-5">
        <p>Subsystems Status:</p>
        <ul class="list-group">
          <li class="list-group-item d-flex align-items-center">Database: <StatusIndicator :on="subsystems.database"></StatusIndicator></li>
          <li class="list-group-item d-flex align-items-center">Connection Monitor: <StatusIndicator :on="subsystems.internet_monitor"></StatusIndicator></li>
          <li class="list-group-item d-flex align-items-center">Timekeeper: <StatusIndicator :on="subsystems.ntp_timekeeper"></StatusIndicator></li>
          <li class="list-group-item d-flex align-items-center">Active Trader: <StatusIndicator :on="subsystems.strategy"></StatusIndicator></li>
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
    ...mapGetters("public", {
      subsystems: "getSubStatus",
    }),
  },
  methods: {
    ...mapActions("public", {
      actionSubStatus: "getSubsystemStatus",
    }),
    async getAll() {
      this.loading = true;
      await this.actionSubStatus().then(
        (res) => {
          //this.subsystems = res
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
