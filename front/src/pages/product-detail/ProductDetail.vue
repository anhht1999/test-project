<template>
  <section class="section product-detail">
    <div class="details container">
      <div class="left">
        <div class="main">
          <img :src="'http://localhost:3000' + product.url" />
        </div>
      </div>
      <div class="right">
        <h3>{{ product.name }}</h3>
        <div class="price">{{ product.price }}</div>
        <br />
        <button
          class="btn-num-product-down color1 flex-c-m size7 bg8 eff2"
          @click="updateProductQuantity(product.quantity - 1)"
        >
          <i class="fs-12 fa fa-minus" aria-hidden="true"></i>
        </button>

        <input
          class="m-text18 t-center num-product"
          type="number"
          name="num-product"
          :value="product.quantity"
          @input="updateProductQuantity($event.target.value)"
        />

        <button
          class="btn-num-product-up color1 flex-c-m size7 bg8 eff2"
          @click="updateProductQuantity(product.quantity + 1)"
        >
          <i class="fs-12 fa fa-plus" aria-hidden="true"></i>
        </button>
        <br />
        <br />
        <button
          class="flex-c-m sizefull bg1 bo-rad-23 hov1 s-text1 trans-0-4"
          @click="addCart"
        >
          Add to Cart
        </button>
        <br />
        <br />
        <h3>Product Detail</h3>
        <p>
          {{ product.description }}
        </p>
      </div>
    </div>
  </section>
</template>

<script>
import { mapState, mapMutations } from "vuex";

export default {
  name: "ProductDetail",

  computed: mapState("products", ["products", "product"]),
  created() {
    // Get product detail
    this.$store.dispatch("products/getProductById", this.$route.params.id);
    // Get related products
    // this.$store.dispatch("products/getFeaturedProducts");
  },
  async beforeRouteUpdate(to) {
    this.$store.dispatch("products/getProductById", to.params.id);
  },
  methods: {
    addCart() {
      // this.$store.dispatch("cart/addProductToCart", this.product);
      this.addProductToCart(this.product);
      // console.log(this.$store.dispatch("cart/addProductToCart", this.product))
    },
    ...mapMutations("products", ["updateProductQuantity"]),
    ...mapMutations("cart", ["addProductToCart"]),
  },
};
</script>
