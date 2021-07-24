import axios from "axios";

const state = () => ({
  orders: [],
  order_items: [],
  subTotal: 0,
  revenues: [],
  date: [],
  rev: [],
  totals: [],
});

const getters = {

  subTotal(state) {
    return state.order_items.reduce(
      (totalPrice, product) => totalPrice + product.quantity * product.price, 0);
  },
};

const actions = {

  async getOrders({ commit }) {
    const res = await axios.get("http://localhost:8000/orders", { withCredentials: true });
    commit("setOrders", res);
  },

  async GetOrderByID({ commit }, id) {
    try {
      const res = await axios.get(`http://localhost:8000/orders/${id}`)
      commit("setOderById", res)
    } catch (e) {
      console.log(e)
    }
  },

  async updateStatus({ commit }, status) {
    
    try {
      const tes = await axios.put(`http://localhost:8000/orders`, status)
      console.log(tes);
      const res = await axios.get("http://localhost:8000/orders", { withCredentials: true });
      commit("setOrders", res);
    } catch (e) {
      console.log(e)
    }
  },

  async getRevenues({ commit }) {
    const res = await axios.get(`http://localhost:8000/revenues`, { withCredentials: true })
    commit("setRevenues", res)
  }
};

const mutations = {

  setOrders(state, res) {
    state.orders = res.data;
  },

  setOderById(state, res) {
    state.order_items = res.data;
  },

  setRevenues(state, res){
    state.revenues = res.data
    for (let i = 0; i < state.revenues.length; i++ ){
        state.date.push(state.revenues[i].date)
        state.rev.push(state.revenues[i].revenue)
    }
}
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};