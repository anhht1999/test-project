<template>
  <div class="container">
    <div class="py-5 text-center"></div>
    <div class="row">
      <div class="col-md-4 order-md-2 mb-4">
        <h4 class="d-flex justify-content-between align-items-center mb-3">
          <span class="text-muted">Your cart</span>
          <span class="badge badge-secondary badge-pill">{{ totalItems }}</span>
        </h4>
        <ul class="list-group mb-3" v-for="cart in carts" :key="cart.id">
          <li
            class="list-group-item d-flex justify-content-between lh-condensed"
          >
            <div>
              <h6 class="my-0">{{ cart.name }}</h6>
              <small class="text-muted">{{ cart.description }}</small>
            </div>
            <span class="text-muted">${{ total(cart) }}</span>
          </li>
          <li class="list-group-item d-flex justify-content-between">
            <span>Total (USD)</span>
            <strong>${{ subTotal }}</strong>
          </li>
        </ul>
      </div>
      <div class="col-md-8 order-md-1">
        <h4 class="mb-3">Billing address</h4>
        <form class="needs-validation" novalidate >
          <div class="row">
            <div class="col-md-6 mb-3">
              <label for="firstName">First name</label>
              <input
                :value="user.first_name"
                disabled
                type="text"
                class="form-control"
                id="firstName"
                placeholder=""
              />
              <!-- <div class="invalid-feedback">Valid first name is required.</div> -->
            </div>
            <div class="col-md-6 mb-3">
              <label for="lastName">Last name</label>
              <input
                :value="user.last_name"
                disabled
                type="text"
                class="form-control"
                id="lastName"
                placeholder=""
                required
              />
              <!-- <div class="invalid-feedback">Valid last name is required.</div> -->
            </div>
          </div>

          <div class="mb-3">
            <label for="email">Email</label>
            <input
              :value="user.email"
              disabled
              type="email"
              class="form-control"
              id="email"
            />
            <div class="invalid-feedback">
              Please enter a valid email address for shipping updates.
            </div>
          </div>

          <div class="mb-3">
            <label for="address">Address</label>
            <input
              v-model="address"
              type="text"
              class="form-control"
              id="address"
              placeholder="1234 Main St"
              required
            />
            <div class="invalid-feedback">
              Please enter your shipping address.
            </div>
          </div>

          <div class="mb-3">
            <label for="phone number ">Phone Number</label>
            <input
              v-model="phone_number"
              type="text"
              class="form-control"
              id="phone_number"
              placeholder="Phone Number"
              required
            />
            <div class="invalid-feedback">Please enter your phome number.</div>
          </div>

          <div class="mb-3">
            <label for="order">Order Note</label>
            <div class="input-group">
              <input
                v-model="order_notes"
                type="text"
                class="form-control"
                id="order_notes"
                placeholder="Order Note"
              />
            </div>
          </div>
          <hr class="mb-4" />
          <button
            @click ="orders"
            class="btn btn-primary btn-lg btn-block btn-color"
            type="button"
          >
            Continue to checkout
          </button>

          <router-link
            to="/cart"
            class="btn btn-primary btn-lg btn-block btn-color btn-color"
          >
            Back to Cart
          </router-link>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapGetters} from "vuex";

export default {
  name: "paymentPage",
  components: {},
  data() {
    return {
      order_item:{
        product_id : "",
        quantity : "",
      },
      order: {
        user_id: "",
        phone_number: "",
        address: "",
        order_notes: "",
        total_price: "",
        order_items: [],
      },
    };
  },
  methods: {
    orders() {
     this.order_items = this.carts.map( cart =>
        this.order_item = {
          product_id : cart.id,
          quantity : cart.quantity,
        });
      this.$store
        .dispatch("cart/orders", {
          user_id: this.user.id,
          phone_number: this.phone_number,
          address: this.address,
          order_notes: this.order_notes,
          total_price: this.subTotal,
          order_items: this.order_items,
        })
        .then( () => {
            this.$router.push("/confirm");
        });
    },
  },
  computed: {
    ...mapState("cart", ["carts","isOrderSuccess","order"]),
    ...mapState("users", ["user"]),
    ...mapGetters("cart", ["totalItems", "subTotal", "total"]),
  },
};
</script>

<style scoped>
.bd-placeholder-img {
  font-size: 1.125rem;
  text-anchor: middle;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
.form-control {
  border: 1px solid #ced4da !important;
}

.btn-color {
  color: azure;
}
@media (min-width: 768px) {
  .bd-placeholder-img-lg {
    font-size: 3.5rem;
  }
}
</style>
