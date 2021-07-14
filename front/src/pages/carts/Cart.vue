<template>
  <div class="breadcrumbs">
    <div class="container">
      <div class="row">
        <div class="col-12">
          <div class="bread-inner">
            <ul class="bread-list">
              <li>
                <a href="index1.html">Home<i class="ti-arrow-right"></i></a>
              </li>
              <li class="active"><a href="blog-single.html">Cart</a></li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="shopping-cart section">
    <div class="container">
      <div class="row">
        <div class="col-12">
          <!-- Shopping Summery -->
          <table class="table shopping-summery">
            <thead>
              <tr class="main-hading">
                <th>PRODUCT</th>
                <th>NAME</th>
                <th class="text-center">UNIT PRICE</th>
                <th class="text-center">QUANTITY</th>
                <th class="text-center">TOTAL</th>
                <th class="text-center"></th>
              </tr>
            </thead>
            <tbody v-for="cart in carts" :key="cart.id">
              <tr>
                <td class="image" data-title="No">
                  <img :src="'http://localhost:3000' + cart.Url" alt="#" />
                </td>
                <td class="product-des" data-title="Description">
                  <p class="product-name">
                    <a href="#">{{ cart.name }}</a>
                  </p>
                  <p class="product-des">{{ cart.description }}</p>
                </td>
                <td class="price" data-title="Price">
                  <span>${{ cart.price }}</span>
                </td>
                <td class="qty" data-title="Qty">
                  <!-- Input Order -->
                  <div class="input-group">
                    <div class="button minus">
                      <button
                        type="button"
                        class="btn btn-primary btn-number"
                        data-type="minus"
                        @click="
                          updateProductQuantity({
                            productId: cart.id,
                            value: cart.quantity - 1,
                          })
                        "
                      >
                        <i class="ti-minus"></i>
                      </button>
                    </div>
                    <input
                      type="number"
                      class="input-number"
                      data-min="1"
                      :value="cart.quantity"
                      @input="
                        updateProductQuantity({
                          productId: cart.id,
                          value: $event.target.value,
                        })
                      "
                    />
                    <div class="button plus">
                      <button
                        type="button"
                        class="btn btn-primary btn-number"
                        data-type="plus"
                        @click="
                          updateProductQuantity({
                            productId: cart.id,
                            value: cart.quantity + 1,
                          })
                        "
                      >
                        <i class="ti-plus"></i>
                      </button>
                    </div>
                  </div>
                  <!--/ End Input Order -->
                </td>
                <td class="total" data-title="Total">
                  <span>${{ total }}</span>
                </td>
                <td class="action" data-title="Remove">
                  <a
                    @click="
                      removeProductToCart({
                        productId: cart.id,
                      })
                    "
                    ><i class="ti-trash remove-icon"></i
                  ></a>
                </td>
              </tr>
            </tbody>
          </table>
          <!--/ End Shopping Summery -->
        </div>
      </div>
      <div class="row">
        <div class="col-12">
          <!-- Total Amount -->
          <div class="total-amount">
            <div class="row">
              <div class="col-lg-8 col-md-5 col-12"></div>
              <div class="col-lg-4 col-md-7 col-12">
                <div class="right">
                  <ul>
                    <li class="last">
                      You Pay<span>${{ subTotal }}</span>
                    </li>
                  </ul>
                  <div class="button5">
                    <a href="#" class="btn">Process To Checkout</a>
                    <a href="#" class="btn">Continue shopping</a>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!--/ End Total Amount -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapMutations, mapGetters } from "vuex";
export default {
  name: "Cart",

  computed: {
    ...mapState("cart", ["carts", "isLoading"]),
    ...mapGetters("cart", ["subTotal", "total"]),
  },
  methods: {
    // currency,
    ...mapMutations("cart", ["updateProductQuantity"]),
    ...mapMutations("cart", ["removeProductToCart"]),
  },
};
</script>
