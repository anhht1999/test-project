import axios from "axios";
import { API_DOMAIN } from "@/config";

export default {
  async getProducts({ page, limit, sort, order, search, categoryId }) {
    let filterCategory = "";
    if (categoryId) filterCategory = "&category=" + categoryId;

    return axios
      .get(
        `${API_DOMAIN}/products?page=${page}&limit=${limit}&order=${order}&sort=${sort}&search=${search}${filterCategory}`
      )
      .then((response) => {
        return response.data
      });
  },

  async getCategories() {
    return axios.get(`${API_DOMAIN}/categories`).then((response) => {
      return response.data;
    });

  },

  async getProductById(productId) {
    return axios.get(`${API_DOMAIN}/products/${productId}`).then((response) => {
      return response.data;
    });
  },

};
