<template>
  <div class="container py-4">
    <div class="p-5 mb-4 bg-light rounded-3">
      <div class="container-fluid py-5">
        <h3>{{ content }}</h3>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from "vuex";
export default {
  name: "Home",
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
      //this.loading = true;
      await this.actionGetAll().then(
          (res) => {
            //this.content = res
          },
          (error) => {
            //this.loading = false;
            //this.message = error.toString() + " | " + error.response.data.error;
          }
      );
    },
  },
  beforeMount() {
    this.getAll();
  },
};
</script>
