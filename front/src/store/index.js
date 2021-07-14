import { createStore } from "vuex";
import products from "./modules/products.module";
import cart from "./modules/carts.module";
import users from "./modules/users.module";

const debug = process.env.NODE_ENV !== "production";

const store = createStore({
  strict: debug,
  modules: {
    users,
    products,
    cart,
  },
});

export default store;
