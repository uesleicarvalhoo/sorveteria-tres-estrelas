<template>
  <form-sale v-on:submit="registerSale"></form-sale>
</template>

<script>
import FormSale from "./components/FormSale.vue"
import { dispatchRegisterSale } from "../controller/sales"
import router from "../routers"

export default {
  name: "RegisterSale",
  components: {
    FormSale
  },
  methods: {
    async registerSale (sale) {
      await dispatchRegisterSale(sale)
      router.push({ name: "view-sales" })
    }
  }
}
</script>
