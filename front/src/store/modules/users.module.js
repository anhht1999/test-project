import axios from "axios";

const state = () => ({
  user: {},
  isLoginSuccess: false,
  loginMessage: "",
  isRegisterSuccess: false,
  registerMessage: "",
});

const getters = {
  isAuthenticated: state => !!state.user
};

const actions = {

  //register
  async register({ commit }, user) {
    try {
      const res = await axios.post('http://localhost:8000/register',user)
      console.log("register success" + res)
      commit("setRegisterSuccess", true);
      commit("setRegisterMessage", "");
    } catch (e) {
      console.log(e)
      commit("setRegisterMessage", e.message === "Request failed with status code 400" ? "Email already in use" : "Create new account is failed");
      commit("setRegisterSuccess", false);
    }
  },

  //login

  async login({ commit }, user) {
    try {
      const res = await axios.post("http://localhost:8000/login", user, {withCredentials: true})
      localStorage.setItem("User", JSON.stringify(res.data.User))
      commit("setLoginSuccess", true);
      commit("setLoginMessage", "");
      commit("setUser", res.data.User );
    } catch (error) {
      console.log(error)
      commit("setLoginSuccess", false);
      commit(
          "setLoginMessage", "User name or password is wrong!"
      );
    }
  },

  //logout
  async logout({ commit}){
    await axios.post("http://localhost:8000/logout","hello", {withCredentials: true});
    await commit("logout")
    
},

};

const mutations = {
  updateUser(state, user) {
    state.user = user;
    state.isLoginSuccess = true;
  },
  setUser(state, user) {
    state.user = user;
  },

  setLoginSuccess(state, status) {
    console.log(status);
    state.isLoginSuccess = status;
  },

  setLoginMessage(state, message) {
    state.loginMessage = message;
  },

  setRegisterSuccess(state, status) {
    state.isRegisterSuccess = status;
  },

  setRegisterMessage(state, message) {
    state.registerMessage = message;
  },

  logout(state) {
    state.user = {};
    state.isLoginSuccess = false;
    state.isRegisterSuccess = false;
    localStorage.removeItem("User");
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
