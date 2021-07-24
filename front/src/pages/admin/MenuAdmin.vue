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
        <div class="tab-pane" id="messages" role="tabpanel">
          <BarChart :date= "date" :rev="rev"/>
        </div>
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
                <th scope="col">Crated At</th>
                <th scope="col">Name</th>
                <th scope="col">Email</th>
                <th scope="col">Address</th>
                <th scope="col">Phone number</th>
                <th scope="col">Order Note</th>
                <th scope="col">Total</th>
                <th scope="col">Status</th>
                <th scope="col">Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="order in orders" :key="order.id">
                <th scope="row">{{ order.created_at }}</th>
                <td>{{ order.first_name }} {{ order.last_name }}</td>
                <td>{{ order.email }}</td>
                <td>{{ order.address }}</td>
                <td>{{ order.phone_number }}</td>
                <td>{{ order.order_note }}</td>
                <td>{{ order.total_price }}$</td>
                <td v-if="order.status === true">Sucess</td>
                <td v-else>Pending</td>
                <td>
                  <a @click="detailModal(order.id)"
                    ><i class="ti-eye edit-icon"></i
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
                <!-- <input
                  accept="image/*"
                  type="file"
                  class="form-control-file"
                  id="image"
                  @change="onFileChange"
                /> -->
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

    <!-- detail order Modal -->
    <div
      :class="{ show: isOpenModalDetail, 'show-edit': isOpenModalDetail }"
      class="modal fade"
      id="detail"
    >
      <div class="modal-dialog modal-dialog-centered add-modal" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h4 class="modal-title" id="exampleModalLongTitle">Detail Order</h4>

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
              <table class="table">
                <thead>
                  <tr>
                    <th scope="col">Name</th>
                    <th scope="col">Price</th>
                    <th scope="col">Quantity</th>
                    <th scope="col">Total</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="od in order_items" :key="od.id">
                    <th scope="row">{{ od.name }}</th>
                    <td>{{ od.price }}$</td>
                    <td>{{ od.quantity }}</td>
                    <td>{{ od.quantity * od.price }}$</td>
                  </tr>
                </tbody>
              </table>
              <div class="form-group">
                <label for="exampleFormControlFile1">Total Order: </label>
                <span> {{ subTotal }}$</span>
              </div>
              <div class="form-group">
                <label for="exampleFormControlFile1">Status</label>
                <select  v-model="status">
                  <option value="0">Pending</option>
                  <option value="1">Sucess</option>
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
              @click="updateStatus()"
            >
              Update status
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- End detail order Modal -->
  </div>
</template>

<script>
import { mapState, mapGetters, mapActions } from "vuex";
import Pagination from "@/components/Pagination.vue";
import api from "@/service/upload.service";
import BarChart from "@/components/BarChart.vue";

export default {
  name: "MenuAdmin",
  components: {
    Pagination,
    BarChart,
  },
  data() {
    return {
      label : [],
      status: "",
      isOpenModalDetail: false,
      isOpenModal: false,
      productId: 0,
      orderId: 0,
      editProduct: "",
      url: "",
      img: [],
      id: "",
      name: "",
      description: "",
      price: "",
      trade_mark: "",
      category_id: "",
      image: [],
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
    ...mapState("orders", [
      "orders",
      "order_items",
      "revenues",
      "rev",
      "date",
      "isLoading",
    ]),
    ...mapGetters("products", ["itemStartIndex", "itemEndIndex"]),
    ...mapGetters("orders", ["subTotal"]),
  },
  created() {
    this.$store.dispatch("products/getProducts", {});
    this.$store.dispatch("products/getCategories");
    this.$store.dispatch("orders/getOrders");
    this.$store.dispatch("orders/getRevenues", {});
  },
  methods: {
    onFileChange(e) {
      const file = e.target.files[0];
      this.url = URL.createObjectURL(file);

      let formData = new FormData();
      if (file != "") {
        formData.append("file", file);
        api
          .uploadImages(formData)
          .then((response) => {
            this.image.push(response);
          })
          .catch(function () {
            console.log("FAILURE!!");
          });
      }
    },
    changePage(pageIndex) {
      this.$store.dispatch("products/getProducts", { pageIndex });
    },
    OpenModal(product) {
      this.isOpenModal = true;
      this.productId = product.id;
      this.editProduct = product;
    },
    closeModal() {
      this.isOpenModal = false;
      this.isOpenModalDetail = false;
    },
    addProduct() {
      for (let i = 0; i < this.image.length; i++) {
        this.img.push({ url: this.image[i] });
      }
      this.$store.dispatch("products/addProduct", {
        name: this.name,
        description: this.description,
        price: parseInt(this.price),
        category_id: this.categoryId,
        trade_mark: this.trade_mark,
        image: this.img,
      });
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
      this.isOpenModal = false;
    },
    detailModal(id) {
      console.log(id);
      this.isOpenModalDetail = true;
      this.orderId = id;
      this.$store.dispatch("orders/GetOrderByID", this.orderId);
    },
    updateStatus() {
      this.$store.dispatch("orders/updateStatus", {
        id: this.orderId,
        status: this.status,
      });
       this.isOpenModalDetail = false;
    },
    ...mapActions("products", [
      "getProducts",
      "deleteProduct",
      "getCategories",
    ]),
    ...mapActions("orders", ["getOrders"]),
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
.dashboard {
    display: flex;
    /* flex-direction: ; */
    padding-top: 110px;
}
.chart{
    width: 50%;
    height: 100%;
}
</style>