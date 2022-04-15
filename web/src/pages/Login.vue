<template>

  <div id="qst-login-form">

    <div class="text-center">
      <h3>QuantstopTerminal</h3>
      <i><small>{{version.version}}</small></i>
    </div>


    <Form @submit="login" :validation-schema="schema">
      <div class="form-group">
        <label for="username">Username</label>
        <Field name="username" type="text" class="form-control" />
        <ErrorMessage name="username" class="error-feedback" />
      </div>
      <div class="form-group">
        <label for="password">Password</label>
        <Field name="password" type="password" class="form-control" />
        <ErrorMessage name="password" class="error-feedback" />
      </div>

      <div class="form-group d-flex justify-content-center" id="qst-login-button-container">
        <button class="btn btn-primary btn-block" :disabled="loading">
          <span v-show="loading" class="spinner-border spinner-border-sm"></span>
          <span>Login</span>
        </button>
      </div>

      <div class="form-group">
        <div v-if="message" class="alert alert-danger" role="alert">
          {{ message }}
        </div>
      </div>
    </Form>
  </div>

</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";
import { mapActions, mapGetters, } from "vuex";
export default {
  name: "Login",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    const schema = yup.object().shape({
      username: yup.string().required("Username is required!"),
      password: yup.string().required("Password is required!"),
    });

    return {
      loading: false,
      message: "",
      schema,
    };
  },
  computed: {
    ...mapGetters("auth", {
      getLoginApiStatus: "getLoginApiStatus",
    }),
    ...mapGetters("public", {
      version: "getVersion",
    }),
  },
  methods: {
    ...mapActions("auth", {
      actionLoginApi: "loginApi",
    }),
    ...mapActions("public", {
      actionGetVersion: "getVersion",
    }),
    async getVersion() {
      await this.actionGetVersion().then(
          () => {
          },
          (error) => {
            this.message = error.toString() + " | " + error.response.status;
          }
      );
    },
    async login(user) {
      this.loading = true;
      const payload = {
        username: user.username,
        password: user.password
      };
      await this.actionLoginApi(payload).then(
        () => {
          this.loading = false;
          this.$router.push("/home");
        },
        (error) => {
          this.loading = false;
          this.message = error.toString() + " | " + error.response.data.error;
        }
      );
    },

  },
  beforeMount() {
    this.getVersion();
  },
};
</script>

<style scoped>
label {
  font-weight: normal;
  font-size: 12px;
  display: inline-block;
  margin-top: 10px;
  line-height: 1.2857142857rem;
}
#qst-login-form {
  max-width: 400px;
  margin-top: 80px;
  margin-left: auto;
  margin-right: auto;
  padding: 20px;
  background-color: var(--background-color-primary) !important;
}
.error-feedback {
  color: red;
}

#qst-login-button-container {
  padding-top: 30px;
  padding-bottom: 20px;
}
</style>
