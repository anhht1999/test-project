<template>
  <section class="section product-detail">

    <div class="breadcrumbs">
      <div class="container">
        <div class="row">
          <div class="col-12">
            <div class="bread-inner">
              <ul class="bread-list">
                <li>
                  <router-link to="/">Home<i class="ti-arrow-right"></i></router-link>
                </li>
                <li class="active"><router-link to="/product">Product<i class="ti-arrow-right"></i></router-link></li>
                <li class="active"><a href="#">{{product.name}}</a></li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="details container">
      <div class="left">
        <div class="main">
          <img
            v-if="product.image"
            :src="'http://localhost:8000/' + product.image[0].url"
          />
        </div>
        <div class="thumbnails">
          <div class="thumbnail" v-for="image in product.image" :key="image.id">
            <img
              :src="'http://localhost:8000/' + image.url"
            />
          </div>
        </div>
      </div>
      <div class="right">
        <h3>{{ product.name }}</h3>
        <div class="price">$ {{ product.price }}</div>
        <br />
        <!-- <button
          class="btn-num-product-down color1 flex-c-m size7 bg8 eff2"
          @click="updateProductQuantity(product.quantity - 1)"
        >
          <i class="fs-12 fa fa-minus" aria-hidden="true"></i>
        </button> -->

        <input
          class="m-text18 t-center num-product"
          type="number"
          name="num-product"
          value="1"
        />
        <!-- <button
          class="btn-num-product-up color1 flex-c-m size7 bg8 eff2"
          @click="updateProductQuantity(product.quantity + 1)"
        >
          <i class="fs-12 fa fa-plus" aria-hidden="true"></i>
        </button> -->
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

    <div class="related-product">
      <div class="title">
        <h2>Related Products</h2>
      </div>

      <div class="container">
        <div class="row">
          <div
            class="col-lg-3 col-md-3 col-12"
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
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { mapState, mapMutations } from "vuex";

export default {
  name: "ProductDetail",
  computed: {
    ...mapState("products", ["products", "product"]),
  },
  data() {
    return {
      quantity: 1
    };
  },
  created() {
    this.$store.dispatch("products/getProductById", this.$route.params.id);
  },
  async beforeRouteUpdate(to) {
    this.$store.dispatch("products/getProductById", to.params.id);
  },
  methods: {
    addCart() {
      this.addProductToCart(this.product);
    },
    ...mapMutations("products", ["updateProductQuantity"]),
    ...mapMutations("cart", ["addProductToCart"]),
  },
};
</script>

<style scoped>
.m-text18 {
  width: 90px;
}

button {
  padding: 10px 20px 10px 20px;
}

.title {
  text-align: center;
  margin-top: 30px;
}
.thumbnails{
  margin-top: 20px;
}
.details{
  margin-top: 20px;
}
</style>

