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
            class="list-group-item d-flex justify-content-between lh-condensed"
          >
            <div>
              <h6 class="my-0">Product name</h6>
              <small class="text-muted">Brief description</small>
            </div>
            <span class="text-muted">$12</span>
          </li>
          <li class="list-group-item d-flex justify-content-between">
            <span>Total (USD)</span>
            <strong>$20</strong>
          </li>
        </ul>
      </div>
      <div class="col-md-8 order-md-1">
        <h4 class="mb-3">Confirm Information</h4>
        <div class="row mb-2">
          <div class="col-3 mt-1"><h6>First name :</h6></div>
          <div class="col-9"><p>anh</p></div>
        </div>
        <div class="row mb-2">
          <div class="col-3 mt-1"><h6>Last name :</h6></div>
          <div class="col-9"><p class="text-secondary">le</p></div>
        </div>
        <div class="row mb-2">
          <div class="col-3 mt-1"><h6>Username :</h6></div>
          <div class="col-9"><p class="text-secondary">TRONG NHAN</p></div>
        </div>
        <div class="row mb-2">
          <div class="col-3 mt-1"><h6>Email :</h6></div>
          <div class="col-9"><p class="text-secondary">abc@example.com</p></div>
        </div>
        <div class="row mb-2">
          <div class="col-3 mt-1">
            <h6>Address :</h6>
          </div>
          <div class="col-9">
            <p class="text-secondary">abc street xyc city</p>
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
export default {
  name: "ConfirmCheckout",
};
</script>

<style scoped></style>
