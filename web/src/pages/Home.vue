<template>
  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    <h2>Dashboard</h2>
    <div class="btn-toolbar mb-2 mb-md-0">
      <div class="btn-group me-2">
        <button type="button" class="btn btn-sm btn-outline-secondary">Share</button>
        <button type="button" class="btn btn-sm btn-outline-secondary">Export</button>
      </div>
      <button type="button" class="btn btn-sm btn-outline-secondary dropdown-toggle">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-calendar" aria-hidden="true"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
        This week
      </button>
    </div>
  </div>

  <div class="container-fluid">
    <div class="row">
      <div class="col-lg-4">

        <div class="rounded-3 card">
          <div class="card-header">
            Subsystems Status
          </div>
          <ul class="list-group list-group-flush">
            <li class="list-group-item d-flex justify-content-between align-items-start">
              <div class=" me-auto">Database</div>
              <StatusIndicator :value="subsystems.database" :height=24 :width=24></StatusIndicator>
            </li>
            <li class="list-group-item d-flex align-items-center">
              <div class=" me-auto">Connection Monitor</div>
              <StatusIndicator :value="subsystems.internet_monitor" :height=24 :width=24></StatusIndicator>
            </li>
            <li class="list-group-item d-flex align-items-center">
              <div class=" me-auto">Timekeeper</div>
              <StatusIndicator :value="subsystems.ntp_timekeeper" :height=24 :width=24></StatusIndicator>
            </li>
            <li class="list-group-item d-flex align-items-center">
              <div class=" me-auto">Active Trader</div>
              <StatusIndicator :value="subsystems.strategy" :height=24 :width=24></StatusIndicator>
            </li>
          </ul>
        </div>


      </div>
      <div class="col-lg-4">

        <div class="rounded-3 card">
          <div class="card-header">
            System Uptime
          </div>
          <div class="d-flex flex-wrap justify-content-center mt-2">
            <h2>{{uptime}}</h2>
<!--            <h3>{{hours}}:</h3>
            <h3>{{mins}}:</h3>
            <h3>{{secs}}</h3>-->
          </div>
        </div>

      </div>
      <div class="col-lg-4">

      </div>
    </div>

  </div>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from "vuex";
import StatusIndicator from "../components/StatusIndicator";
import jwtInterceptor from "../shared/jwt.interceptor";
export default {
  name: "Home",
  components: {StatusIndicator},
  computed: {
    ...mapGetters("public", {
      subsystems: "getSubStatus",
    }),
    ...mapGetters("auth", {
      getUserProfile: "getUserProfile",
    }),
  },
  data () {
    return {
      /*hours: 0,
      mins: 0,
      secs: 0,*/
      uptime: "0d 0h 0m 0s"
    }
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
    async getUptime() {
      const response = await jwtInterceptor.get("/api/uptime", {
        withCredentials: true,
        credentials: "include",
        headers: {
          'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
        },
      });
      if (response && response.data) {

        this.uptime = response.data
        /*if (response.data.includes("h")) {
          this.hours = response.data.substring(0, response.data.indexOf("h"));
        } else {
          this.hours = "00";
        }


        this.mins = response.data.substring(response.data.indexOf("h"), response.data.indexOf("m"));
        this.secs = response.data.substring(response.data.indexOf("m") +1, response.data.indexOf("s"));
        this.secs = this.secs.substring(0, this.secs.indexOf("."))*/
      }
    },
  },
  beforeMount() {
    this.getAll();
    this.getUptime();
  },
};
</script>
