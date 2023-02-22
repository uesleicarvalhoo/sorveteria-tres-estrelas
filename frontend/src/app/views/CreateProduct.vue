<template>
  <form-product v-on:submit="createProduct"></form-product>
</template>

<script>
import FormProduct from "./components/FormProduct.vue"
import { dispatchCreateProduct } from "../controller/products"
import router from "../routers"

export default {
  name: "CreateItemForm",
  components: {
    FormProduct: FormProduct
  },
  methods: {
    async createProduct (item) {
      await dispatchCreateProduct(item)
      router.push({ name: "view-products" })
    }
  }
}
</script>
