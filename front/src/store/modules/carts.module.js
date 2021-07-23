import axios from "axios";

const state = () => ({
  order: [],
  carts: [],
  addToCartResult: "",
  totalItems: 0,
  subTotal: 0,
  isOrderSuccess: false,
});

const getters = {
  totalItems(state) {
    return state.carts.reduce((total, product) => total + product.quantity, 0);
  },

  subTotal(state) {
    return state.carts.reduce(
      (totalPrice, product) => totalPrice + product.quantity * product.price,
      0
    );
  },

  total() {
    return (product) => product.price * product.quantity;
  },
};

const actions = {


  async orders({ commit }, order) {
    try {
      const res = await axios.post("http://127.0.0.1:8000/orders", order);
      console.log(res);
      commit("setOrderSuccess", true);
      // commit("setOrder", order);
    } catch (e) {
      console.log(e);
      commit("setOrderSuccess", false);
    }
  },
};

const mutations = {

  setOrder(state, carts) {
    state.order.push(carts)
  },

  setOrderSuccess(state, status) {
    state.isOrderSuccess = status;
  },

  setProducts(state, carts) {
    state.carts = carts.map((product) => {
      product.totalPrice = product.quantity * product.price;
      return product;
    });
  },

  updateProductQuantity(state, { productId, value }) {
    const product = state.carts.find((p) => p.id === productId);

    value = Number(value);

    if (value > 1) {
      product.quantity = value;
    } else {
      product.quantity = 1;
    }

    product.totalPrice = product.price * product.quantity;
  },

  addProductToCart(state, product) {
    const isExists = state.carts.find((p) => p.id === product.id);

    if (isExists) {
      product.quantity++;
    } else {
      state.carts.push(product);
    }
  },

  removeProductToCart(state, { productId }) {
    const product = state.carts.find((p) => p.id === productId);
    if (product) {
      state.carts.splice(product, 1);
    }
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
