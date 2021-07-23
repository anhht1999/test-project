<template>
  <!--Create product Modal -->
  <div
    class="modal fade"
    id="exampleModalCenter"
    tabindex="-1"
    role="dialog"
    aria-labelledby="exampleModalCenterTitle"
    aria-hidden="true"
  >
    <div class="modal-dialog modal-dialog-centered add-modal" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title" id="exampleModalLongTitle">Add Product</h4>

          <button
            type="button"
            class="close"
            data-dismiss="modal"
            aria-label="Close"
          >
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
        <div class="modal-body">
          <form enctype="multipart/form-data">
            <div class="form-group">
              <label for="formGroupExampleInput">Name</label>
              <input
                v-model="name"
                type="text"
                class="form-control"
                id="name"
                placeholder="Name"
              />
            </div>
            <div class="form-group">
              <label for="formGroupExampleInput2">Description</label>
              <input
                v-model="description"
                type="text"
                class="form-control"
                id="description"
                placeholder="Description"
              />
            </div>

            <div class="form-group">
              <label for="formGroupExampleInput2">Price</label>
              <input
                v-model="price"
                type="number"
                class="form-control"
                id="price"
                placeholder="Price"
              />
            </div>
            <div class="form-group">
              <label for="formGroupExampleInput2">Trade mark</label>
              <input
                v-model="trade_mark"
                type="text"
                class="form-control"
                id="trade_mark"
                placeholder="Trade mark"
              />
            </div>

            <div class="form-group">
              <label for="exampleFormControlFile1">Image</label>
              <input
                accept="image/*"
                type="file"
                class="form-control-file"
                id="image"
                @change="onFileChange"
              />
              <img id="preview" v-if="url" :src="url" />
            </div>
            <div class="form-group">
              <label for="exampleFormControlFile1">Caterogy</label>
              <select class="form-control form-control-sm" v-model="categoryId">
                <option
                  v-for="cate in categories"
                  :key="cate.id"
                  :value="cate.id"
                >
                  {{ cate.name }}
                </option>
              </select>
            </div>
          </form>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-dismiss="modal">
            Close
          </button>
          <button type="button" class="btn btn-primary" @click="addProduct()">
            Add
          </button>
        </div>
      </div>
    </div>
  </div>
  <!-- End Create product Modal -->
</template>

<script>

import { mapState, mapActions } from "vuex";

export default {
  name: AddProductForm,
  data() {
    return {
      url: "",
      img: {
        url: "",
      },
      add_product: {
        id:"",
        name: "",
        description: "",
        price: "",
        trade_mark: "",
        category_id: "",
        image: [],
      },
    };
  },
    computed: {
    ...mapState("products", [
      "products",
      "categories",
    ]),
  },
  created() {
    this.$store.dispatch("products/getProducts", {});
    this.$store.dispatch("products/getCategories");
  },
  methods: {
    onFileChange(e) {
      const file = e.target.files[0];
      console.log(file);
      this.url = URL.createObjectURL(file);
    },

     addProduct() {
      this.img = {
        url: this.url,
      };
      this.$store.dispatch("products/addProduct", {
        name: this.name,
        description: this.description,
        price: parseInt(this.price),
        category_id: this.categoryId,
        trade_mark: this.trade_mark,
        image: [this.img],
      });
      if (this.isProductSuccess) {
        this.message = "Create product successfully.";
      }
    },
  },

  ...mapActions("products", [
      "getProducts",,
      "addProduct"
    ]),
};
</script>

<style>
.add-modal {
  width: 500px;
  height: 300px;
}
.modal-body {
  padding-top: 0px;
}

#preview {
  height: 100px;
  width: 100px;
  margin-top: 5px;
  margin-bottom: 5px;
  line-height: 20px;
}
.form-group {
  margin-bottom: 0.5em;
}
.close {
  color: red;
}
.modal-header {
  padding-bottom: 0px;
}
.modal-footer {
  border: 0px;
  padding-top: 0px;
}
#image {
  line-height: 15px;
}
</style>