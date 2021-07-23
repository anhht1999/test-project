<template>
  <section class="product-area shop-sidebar shop section">
    <div class="breadcrumbs">
      <div class="container">
        <div class="row">
          <div class="col-12">
            <div class="bread-inner">
              <ul class="bread-list">
                <li>
                  <a href="index1.html">Home<i class="ti-arrow-right"></i></a>
                </li>
                <li class="active"><a href="blog-single.html">Product</a></li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
    <br>
    <div class="container">
      <div class="row">
        <LeftBar />
        <div class="col-lg-9 col-md-8 col-12">
          <div class="row">
            <div class="col-12">
              <!-- Shop Top -->
              <div class="shop-top">
                <div class="shop-shorter">
                  <div class="single-shorter">
                    <label>Show :</label>
                    <select @change="sortProducts">
                      <option value="id-desc">Defaut</option>
                      <option value="price-desc">Price: Low to High</option>
                      <option value="price-asc">Price: High to Low</option>
                    </select>
                  </div>
                </div>
              </div>
              <!--/ End Shop Top -->
            </div>
          </div>
          <div class="row" v-if="!isLoading">
            <div
              class="col-lg-4 col-md-6 col-12"
              v-for="product in products"
              :key="product.id"
            >
              <div class="single-product">
                <div class="product-img">
                  <router-link :to="'/product/' + product.id">
                    <img
                      class="default-img"
                      :src="'http://127.0.0.1:8000/' + product.image[0].url"
                    />
                  </router-link>
                  <div class="button-head">
                    <div class="product-action">
                      <a
                        data-toggle="modal"
                        data-target="#exampleModal"
                        title="Quick View"
                        href="#"
                        ><i class="ti-eye"></i><span>Quick Shop</span></a
                      >
                    </div>
                    <div class="product-action-2">
                      <a title="Add to cart" @click="addToCart(product)"
                        >Add to cart</a
                      >
                    </div>
                  </div>
                </div>
                <div class="product-content">
                  <h3>
                    <router-link :to="'/product/' + product.id">
                      {{ product.name }}
                    </router-link>
                  </h3>
                  <div class="product-price">
                    <span>${{ product.price }}</span>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-lg-12 col-md-12 col-12">
              <Pagination
                :length="totalItems"
                :pageSize="limit"
                :pageIndex="pageIndex"
                @change="changePage"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import LeftBar from "./LeftBar.vue";
import { mapState, mapGetters, mapActions, mapMutations } from "vuex";
import Pagination from "@/components/Pagination.vue";

export default {
  name: "ProductsPage",

  components: {
    LeftBar,
    Pagination,
  },
  data() {
    return {
      quantity: 1
    };
  },
  computed: {
    ...mapState("products", [
      "isLoading",
      "products",
      "totalItems",
      "pageIndex",
      "limit",
    ]),
    ...mapGetters("products", [
      "sortDropdownValue",
      "itemStartIndex",
      "itemEndIndex",
    ]),
  },

  created() {
    this.$store.dispatch("products/getProducts", {});
  },

  methods: {
    // currency,
    sortProducts(event) {
      const options = event.target.value.split("-");
      console.log(options);
      const sort = options[1],
        order = options[0];
      console.log(sort, order);
      this.$store.dispatch("products/getProducts", { sort, order });
    },
    changePage(pageIndex) {
      this.$store.dispatch("products/getProducts", { pageIndex });
    },
    addToCart(product) {
      console.log(product);
      this.addProductToCart(product);
    },
    ...mapMutations("products", ["updateProductQuantity"]),
    ...mapActions("products", ["getProducts"]),
    ...mapMutations("cart", ["addProductToCart"]),
  },
};
</script>
