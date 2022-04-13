<template>

  <footer class="footer py-3">
    <div class="container">
      <span class="text-center">
        Version: {{version.version}} <br>
        {{version.copyright}}
      </span>
    </div>
  </footer>

</template>

<script>
import {mapActions, mapGetters} from "vuex";

export default {
  name: "Footer",
  computed: {
    ...mapGetters("public", {
      version: "getVersion",
    }),
  },
  methods: {
    ...mapActions("public", {
      actionGetVersion: "getVersion",
    }),
    async getAll() {
      this.loading = true;
      await this.actionGetVersion().then(
          (res) => {
            //this.version = res
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
}
</script>

<style scoped>

</style>