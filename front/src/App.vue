<template>
<div>
    <MainHeader />

    <router-view
      :products="products" 
    ></router-view>

    <MainFooter/>
  </div>
</template>

<script>
import MainHeader from "./components/MainHeader.vue";
import MainFooter from "./components/MainFooter.vue";
import { mapMutations } from "vuex";

export default {
  name:"App",
  components:{
    MainHeader,
    MainFooter,
  },
  methods:{
    ...mapMutations('users',['setUser'])
  },
  created() {
    const user = localStorage.getItem("User") 
    this.setUser(JSON.parse(user))
  }
}
</script>
