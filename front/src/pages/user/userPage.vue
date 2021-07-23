<template lang="">
    <div>
    <h3 v-if ="isLoginSuccess === true">Hello {{user.last_name}}</h3>
    <section class="bgwhite p-t-66 p-b-60">
      <div class="container">
        <div class="row m-t-30">
          <div class="col-sm-6 col-md-4 col-lg-1 p-b-50">
            <ul>
              <li>
                <router-link to="/user/profile">Profile</router-link>
              </li>
              <li>
                <router-link to="/user/settings">Settings</router-link>
              </li>
              <li>
                <a href="/" @click="logout">Logout</a>
              </li>
            </ul>
          </div>

          <div class="col-sm-6 col-md-8 col-lg-11 p-b-50">
            <router-view/>
          </div>
        </div>
      </div>
    </section>
    </div>
</template>
<script>
import { mapState } from "vuex";

export default {
  name: "UserPage",
  computed: {
    ...mapState("users", ["isLoginSuccess", "user"]),
  },
  methods: {
   logout() {
      this.$store.dispatch("users/logout");
    },
  },
};
</script>
<style lang="">
</style>