<template>
  <div class="container">
    <div class="py-5 text-center">
      <h3>Order Success</h3>
    </div>
    <div class="row">
      <div class="col-md-4 order-md-2 mb-4 text-secondary">
        <h4 class="d-flex justify-content-between align-items-center mb-3">
          <span class="text-muted">Your cart</span>
        </h4>
        <ul class="list-group mb-3">
          <li
            v-for="cart in cart_item" :key="cart.id"
            class="list-group-item d-flex justify-content-between lh-condensed"
          >
            <div>
              <h6 class="my-0">{{cart.name}}</h6>
              <small class="text-muted">{{cart.description}}</small>
            </div>
            <span class="text-muted">${{cart.price}}</span>
          </li>
          <li class="list-group-item d-flex justify-content-between">
            <span>Total (USD)</span>
            <strong>${{subTotal}}</strong>
          </li>
        </ul>
      </div>
      <div class="col-md-8 order-md-1">
        <h4 class="mb-3">Confirm Information</h4>
        <div class="row mb-2">
          <div class="col-3 mt-1"><h6>First name :</h6></div>
          <div class="col-9"><p>{{order.first_name}}</p></div>
        </div>
        <div class="row mb-2">
          <div class="col-3 mt-1"><h6>Last name :</h6></div>
          <div class="col-9"><p class="text-secondary">{{order.last_name}}</p></div>
        </div>
        <div class="row mb-2">
          <div class="col-3 mt-1"><h6>Email :</h6></div>
          <div class="col-9"><p class="text-secondary">{{order.email}}</p></div>
        </div>
        <div class="row mb-2">
          <div class="col-3 mt-1">
            <h6>Address :</h6>
          </div>
          <div class="col-9">
            <p class="text-secondary">{{order.address}}</p>
          </div>
        </div>
        <div class="row mb-2">
          <div class="col-3 mt-1">
            <h6>Phone number :</h6>
          </div>
          <div class="col-9">
            <p class="text-secondary">{{order.phone_number}}</p>
          </div>
        </div>
        <hr class="mb-4" />
        <div class="row">
          <div class="col-6">
            <button class="btn btn-primary btn-lg btn-block" type="submit">
              Continue Shopping
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import {mapState} from 'vuex';

export default {
  name: "ConfirmCheckout",
  computed:{
    ...mapState("cart",["order","cart_item"])
  },
  methods:{
    subTotal() {
    return this.cart_item.reduce(
      (totalPrice, product) => totalPrice + product.quantity * product.price,
      0
    );
  },
  }
};
</script>

<style scoped></style>
