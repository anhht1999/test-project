<template>
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
              <select class="form-control form-control-sm" v-model="categoryId">
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
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
    name: EditProductForm,
     computed: {
    ...mapState("products", [
      "products",
      "categories",
    ]),
  },
};
</script>

<style>
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