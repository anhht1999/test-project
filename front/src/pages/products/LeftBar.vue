<template>
  <div class="col-lg-3 col-md-4 col-12">
    <div class="shop-sidebar">
      <div class="single-widget category">
        <h3 class="title">Categories</h3>
        <ul class="categor-list">
          <li>
            <a
              @click="filterProductsByCategory($event, {})"
              :class="{ active: !category.id }"
              >All Product</a
            >
          </li>
          <li
            class="nav-item dropdown"
            v-for="cate in categories"
            :key="cate.id"
          >
            <a
              @click="filterProductsByCategory($event, cate)"
              :class="{ active: cate.id === category.id }"
              >{{ cate.name }}</a
            >
          </li>
        </ul>
      </div>

      <div class="single-widget category">
        <h3 class="title">SEARCH</h3>
        <div class="search-product pos-relative bo4 of-hidden">
          <input
            class="s-text7 size6 p-l-23 p-r-50"
            type="text"
            placeholder="Search Products..."
            v-model="searchKeyword"
            @keypress.enter="searchProducts"
          />

          <button
            class="flex-c-m size5 ab-r-m color2 color0-hov trans-0-4"
            @click="searchProducts"
          >
            <i class="fs-12 fa fa-search" aria-hidden="true"></i>
          </button>
        </div>
      </div>
      <div class="single-widget range">
        <h3 class="title">Shop by Price</h3>
        <ul class="check-box-list">
          <li>
            <label class="checkbox-inline"
              ><input
                name="1-10"
                id="1-10"
                type="checkbox"
                v-model="checkedPrice"
                value="1-10"
                @click="check" />$1 - $10<span></span
            ></label>
          </li>
          <li>
            <label class="checkbox-inline"
              ><input
                name="11-20"
                id="11-20"
                type="checkbox"
                v-model="checkedPrice"
                value="11-20"
                @click="check" />$11 - $20<span></span
            ></label>
          </li>
          <li>
            <label class="checkbox-inline"
              ><input
                name="21-50"
                id="21-50"
                type="checkbox"
                v-model="checkedPrice"
                value="21-50"
                @click="check" />$21 - $50<span></span
            ></label>
          </li>
        </ul>
      </div>
      <div class="single-widget category">
        <h3 class="title">Author</h3>
        <ul class="categor-list">
          <li><a href="#">Forever</a></li>
          <li><a href="#">giordano</a></li>
          <li><a href="#">abercrombie</a></li>
          <li><a href="#">ecko united</a></li>
          <li><a href="#">zara</a></li>
        </ul>
      </div>
      <!--/ End Single Widget -->
    </div>
  </div>
</template>
<script>
import { mapState } from "vuex";

export default {
  name: "LeftBar",
  data() {
    return {
      searchKeyword: "",
      checkedPrice: [],
    };
  },
  computed: {
    ...mapState("products", ["search", "categories", "category"]),
  },
  created() {
    this.searchKeyword = this.search;
    this.$store.dispatch("products/getCategories");
  },
  methods: {
    filterProductsByCategory(event, category) {
      event.preventDefault();
      this.$store.dispatch("products/getProducts", {
        pageIndex: 1,
        category,
      });
    },
    searchProducts() {
      this.$store.dispatch("products/getProducts", {
        pageIndex: 1,
        search: this.searchKeyword,
      });
    },

    // check(){
    //   this.$store.dispatch("products/getProducts", {
    //     pageIndex: 1,
    //     checked: this.checkedPrice,
    //   });
    // }
  },
};
</script>
