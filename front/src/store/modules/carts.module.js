import api from "@/service/cart.service";

const state = () => ({
  carts: [],
  isLoading: false,
  addToCartResult: "",
  totalItems: 0,
  subTotal: 0,
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
  async addProductToCart({ state, commit }, product) {
    const isExists = state.carts.find((p) => p.id === product.id);

    if (isExists) {
      product.quantity = product.quantity + 1;
    } else {
      const newProduct = await api.addProductToCart(product);

      if (newProduct) {
        commit("addProductToCart", newProduct);
      }
    }
  },

  async getProductsInCart({ commit }) {
    commit("setLoading", true);

    const carts = await api.getProductsInCart();

    commit("setProducts", carts);
    commit("setLoading", false);
  },
};

const mutations = {
  setLoading(state, status) {
    state.isLoading = status;
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

  // setAddToCartResult(state, message) {
  //   state.addToCartResult = message;
  // },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
