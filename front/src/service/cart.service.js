import axios from "axios";
import { API_DOMAIN } from "@/config";

export default {
  async addProductToCart(product) {
    return axios.post(`${API_DOMAIN}/orders`, product).then((response) => {
      const product = response.data;
      product.Url = API_DOMAIN + product.url;
      return product;
    });
  },

  async getProductsInCart() {
    return axios.get(`${API_DOMAIN}/orders`).then((response) => {
      const products = response.data.map((product) => {
        product.Url = API_DOMAIN + product.url;
        return product;
      });

      return products;
    });
  },
};
