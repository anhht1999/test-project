<template>
  <!-- content page -->
  <section class="bgwhite p-t-66 p-b-60">
    <h2 class="l-text2 t-center">Login</h2>
    <div class="container">
      <div class="login-form m-auto">
        <Form @submit="login" :validation-schema="schema">
          <p class="m-b-20 text-center text-danger">
            {{ loginMessage }}
          </p>

          <div class="bo4 of-hidden size15 m-b-10">
            <Field
             v-model="user.email"
              name="email"
              type="text"
              placeholder="Email"
              class="sizefull s-text7 p-l-22 p-r-22"
              :disabled="isLoading"
            />
          </div>
            <br />
          <ErrorMessage name="email" class="text-danger m-b-20 d-block" />

          <div class="bo4 of-hidden size15 m-b-10">
            <Field
             v-model="user.password"
              name="password"
              type="password"
              placeholder="Password"
              class="sizefull s-text7 p-l-22 p-r-22"
              :disabled="isLoading"
            />
          </div>

          <ErrorMessage name="password" class="text-danger m-b-20 d-block" />
              <br />
          <div class="w-size25 m-auto">
            <button
              class="flex-c-m size2 bg1 bo-rad-23 hov1 m-text3 trans-0-4 m-t-20 click"
              :class="{ disabled: isLoading }"
              :disabled="isLoading"
            >
              <span v-show="isLoading" data-loader="ball-scale"></span>
              Login
            </button>
          </div>
        </Form>

        <br />
        <hr />

        <p class="text-center">
          <router-link to="/register"
            >Don't have an account yet? Register now!
          </router-link>
        </p>
        <br>
      </div>
    </div>
  </section>
</template>

<script>
import { mapState } from "vuex";
import { Field, Form, ErrorMessage } from "vee-validate";
import * as yup from "yup";
export default {
  name: "LoginPage",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    const schema = yup.object().shape({
      email: yup
        .string()
        .required("Email is required!")
        .email("Email is invalid!"),
      password: yup.string().required("Password is required!"),
    });
    return {
      isLoading: false,
      schema,
      user: {
        email: "",
        password: "",
      },
    };
  },
  computed:{
    ...mapState("users", ["isLoginSuccess", "loginMessage"]),
    ...mapState("cart", ["carts"]),
  },
  created() {
    if (this.isLoginSuccess) {
      console.log(this.isLoginSuccess);
      this.$router.push("/");
    }
  },
  methods: {
    login() {
      this.isLoading = true;
      this.$store
        .dispatch("users/login", {
          email: this.user.email,
          password: this.user.password,
        })
        .then(() => {
          this.isLoading = false;
          if (this.isLoginSuccess && this.carts.length === 0) {
            this.$router.push("/user");
          }
          else{
            this.$router.push("/payment");
          }
        });
    },
  },
};
</script>

<style scoped>
.login-form {
  max-width: 700px;
  text-align: center;
}
button.disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
button.disabled:hover {
  background-color: #222222;
}
.t-center{
   text-align: center;
 }
.l-text2{
    margin: 20px 0 20px 0;
}
.click{
  padding: 15px 30px 15px 30px;
}
</style>
