<template>
  <div class="container rounded bg-white mt-5 mb-5 py-4">
    <div class="row p-5 mb-4 bg-light rounded-3">
      <h3>Email Verification</h3>
      <div
          v-if="message"
          class="alert"
          :class="successful ? 'alert-success' : 'alert-danger'"
      >
        {{ message }}
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters, } from "vuex";
import { useRoute } from 'vue-router'
let code = ""
export default {
  name: 'Verify',
  setup() {
    const route = useRoute();
    code = route.params.code;

  },
  data() {
    return {
      successful: false,
      loading: false,
      message: "",
    };
  },
  computed: {
    ...mapGetters("auth", {
      getVerifyApiStatus: "getVerifyApiStatus",
    }),
  },
  methods: {
    ...mapActions("auth", {
      actionSendVerify: "verifyApi",
    }),
    async sendVerify(code) {
      this.loading = true;
      const payload = {
        code: code,
      };
      await this.actionSendVerify(payload).then(
          () => {
            this.message = "Success! Your email has been verified, you may now login."
            this.successful = true;
            this.loading = false;
          },
          (error) => {
            this.loading = false;
            this.successful = false;
            this.message = error.toString() + " | " + error.response.data.error;
          }
      );
    },
  },
  beforeMount() {
    this.sendVerify(code);
  },
};
</script>