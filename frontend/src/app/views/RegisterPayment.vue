<template>
  <form-balance v-on:submit="createPayment"></form-balance>
</template>

<script>
import FormPayment from "./components/FormPayment.vue"
import { dispatchCreatePayment } from "../controller/payments"
import router from "../routers"

export default {
  name: "CreateBalanceForm",
  components: {
    FormBalance: FormPayment
  },
  methods: {
    async createPayment (payment) {
      await dispatchCreatePayment(payment)
      router.push({ name: "view-payments" })
    }
  }
}
</script>
