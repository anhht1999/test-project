<template>
  <div class="container">
    <br />
    <div class="row">
      <ul
        class="nav nav-tabs col-2 col-md-2 col-lg-2 menu-admin"
        id="myTab"
        role="tablist"
      >
        <li class="nav-item">
          <a
            class="nav-link active"
            data-toggle="tab"
            href="#messages"
            role="tab"
            aria-controls="messages"
            >Dashboard</a
          >
        </li>
        <li class="nav-item">
          <a
            class="nav-link"
            data-toggle="tab"
            href="#home"
            role="tab"
            aria-controls="home"
            >Product</a
          >
        </li>
        <li class="nav-item">
          <a
            class="nav-link"
            data-toggle="tab"
            href="#profile"
            role="tab"
            aria-controls="profile"
            >Order</a
          >
        </li>
        <li class="nav-item">
          <a class="nav-link" role="tab" aria-controls="profile">Logout</a>
        </li>
      </ul>
      <div class="tab-content col-10 col-md-10 col-lg-10">
        <div class="tab-pane" id="messages" role="tabpanel">Chart</div>
        <div class="tab-pane active" id="home" role="tabpanel">
          <!-- Button trigger modal -->
          <button
            type="button"
            class="btn btn-primary"
            data-toggle="modal"
            data-target="#exampleModalCenter"
          >
            Add Product
          </button>
          <br />
          <br />

          <table class="table table-hover">
            <thead>
              <tr>
                <th scope="col-1">Name</th>
                <th scope="col-4">Description</th>
                <th scope="col-1">Price</th>
                <th scope="col-1">Caterogy</th>
                <th scope="col-2">Image</th>
                <th scope="col-2">Trade mark</th>
                <th scope="col-1">Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="product in products" :key="product.id">
                <th scope="row">{{ product.name }}</th>
                <td>{{ product.description }}</td>
                <td>${{ product.price }}</td>
                <td>{{ product.category }}</td>
                <td v-if="product.image">
                  <img
                    class="default-img"
                    :src="'http://127.0.0.1:8000/' + product.image[0].url"
                  />
                </td>
                <td>{{ product.trade_mark }}</td>
                <td class="action">
                  <a @click="OpenModal(product)"
                    ><i class="ti-pencil edit-icon"></i
                  ></a>
                  <a @click="deleteProduct(product.id)"
                    ><i class="ti-trash remove-icon"></i
                  ></a>
                </td>
              </tr>
            </tbody>
          </table>

          <div class="col-lg-12 col-md-12 col-12">
            <Pagination
              :length="totalItems"
              :pageSize="limit"
              :pageIndex="pageIndex"
              @change="changePage"
            />
          </div>
        </div>
        <div class="tab-pane" id="profile" role="tabpanel">
          <table class="table table-hover">
            <thead>
              <tr>
                <th scope="col">order</th>
                <th scope="col">Description</th>
                <th scope="col">Price</th>
                <th scope="col">Caterogy</th>
                <th scope="col">Image</th>
                <th scope="col">Trade mark</th>
                <th scope="col">Action</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <th scope="row">1</th>
                <td>Mark</td>
                <td>Otto</td>
                <td>@mdo</td>
                <td>@mdo</td>
                <td>@mdo</td>
                <td>
                  <a><i class="ti-pencil edit-icon"></i></a>-<a
                    ><i class="ti-trash remove-icon"></i
                  ></a>
                </td>
              </tr>
              <tr>
                <th scope="row">2</th>
                <td>Mark</td>
                <td>Otto</td>
                <td>@mdo</td>
                <td>@mdo</td>
                <td>@mdo</td>
                <td>
                  <a><i class="ti-pencil edit-icon"></i></a>-<a
                    ><i class="ti-trash remove-icon"></i
                  ></a>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

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
                <select
                  class="form-control form-control-sm"
                  v-model="categoryId"
                >
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
            <button
              type="button"
              class="btn btn-secondary"
              data-dismiss="modal"
            >
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

    <!--Edit product Modal -->
    <div
      :class="{ show: isOpenModal, 'show-edit': isOpenModal }"
      class="modal fade"
      id="edit"
      tabindex="-1"
      role="dialog"
      aria-labelledby="edit"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered add-modal" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h4 class="modal-title" id="exampleModalLongTitle">Edit Product</h4>

            <button
              type="button"
              class="close"
              data-dismiss="modal"
              aria-label="Close"
              @click="closeModal()"
            >
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form>
              <div class="form-group">
                <label for="formGroupExampleInput">Name</label>
                <input
                  v-model="editProduct.name"
                  type="text"
                  class="form-control"
                  id="name"
                  placeholder="Name"
                />
              </div>
              <div class="form-group">
                <label for="formGroupExampleInput2">Description</label>
                <input
                  v-model="editProduct.description"
                  type="text"
                  class="form-control"
                  id="description"
                  placeholder="Description"
                />
              </div>

              <div class="form-group">
                <label for="formGroupExampleInput2">Price</label>
                <input
                  v-model="editProduct.price"
                  type="number"
                  class="form-control"
                  id="price"
                  placeholder="Price"
                />
              </div>
              <div class="form-group">
                <label for="formGroupExampleInput2">Trade mark</label>
                <input
                  v-model="editProduct.trade_mark"
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
                <img
                  id="preview"
                  v-if="editProduct.image"
                  :src="'http://127.0.0.1:8000/' + editProduct.image[0].url"
                />
              </div>
              <div class="form-group">
                <label for="exampleFormControlFile1">Caterogy</label>
                <select
                  class="form-control form-control-sm"
                  v-model="categoryId"
                >
                  <option disabled>{{ editProduct.category }}</option>
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
            <button
              type="button"
              class="btn btn-secondary"
              data-dismiss="modal"
              @click="closeModal()"
            >
              Close
            </button>
            <button
              type="button"
              class="btn btn-primary"
              @click="updateProduct()"
            >
              Edit
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- End Modal -->
  </div>
</template>

<script>
import { mapState, mapGetters, mapActions } from "vuex";
import Pagination from "@/components/Pagination.vue";

export default {
  name: "MenuAdmin",
  components: {
    Pagination,
  },
  data() {
    return {
      isOpenModal: false,
      productId: 0,
      editProduct: "",
      url: "",
      img: {
        url: "",
      },
      add_product: {
        id: "",
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
      "totalItems",
      "pageIndex",
      "limit",
      "categories",
      "isProductSuccess",
      "productMessage",
      "isLoading",
    ]),
    ...mapGetters("products", ["itemStartIndex", "itemEndIndex"]),
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
    changePage(pageIndex) {
      this.$store.dispatch("products/getProducts", { pageIndex });
    },
    OpenModal(product) {
      this.isOpenModal = true;
      this.productId = product.id;
      this.editProduct = product;
      // console.log(this.editProduct);
    },
    closeModal() {
      this.isOpenModal = false;
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
        this.$router.push("/admin");
        this.message = "Create product successfully.";
      }
    },

    updateProduct() {
      this.$store.dispatch("products/updateProduct", {
        id: this.editProduct.id,
        name: this.editProduct.name,
        description: this.editProduct.description,
        price: parseInt(this.editProduct.price),
        trade_mark: this.editProduct.trade_mark,
        category_id: this.categoryId,
        image: [this.editProduct.image[0]],
      });
    },

    ...mapActions("products", ["getProducts", "deleteProduct"]),
  },
};
</script>

<style scoped>
.show-edit {
  display: block;
  overflow-x: hidden;
  overflow-y: auto;
}
i {
  font-size: 20px;
}
.edit-icon {
  margin-right: 20px;
}
.default-img {
  height: 150px;
  width: 150px;
}
.container {
  margin-bottom: 20px;
}
.menu-admin {
  margin-right: 0px !important;
  padding-right: 0px !important;
}
.nav.nav-tabs {
  float: left;
  display: block;
  border-bottom: 0;
  margin-right: 0px;
}
.nav-tabs .nav-link {
  border: 1px solid transparent;
  border-top-left-radius: 0.25rem;
  border-top-right-radius: 0.25rem;
  background: #ccc;
}

.nav-tabs .nav-link.active {
  color: #495057;
  border-color: transparent !important;
}
.nav-tabs .nav-link {
  border: 1px solid transparent;
  border-top-left-radius: 0rem !important;
  border-top-right-radius: 0rem !important;
}
.tab-content > .active {
  display: block;
  min-height: 165px;
}
.nav.nav-tabs {
  float: left;
  display: block;
  margin-right: 20px;
  border-bottom: 0;
  border-right: 1px solid transparent;
  padding-right: 15px;
}
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