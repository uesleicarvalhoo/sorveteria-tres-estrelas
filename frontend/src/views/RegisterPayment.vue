<template>
  <form-balance v-on:submit="createPayment"></form-balance>
</template>

<script>
import FormPayment from "./components/FormPayment.vue"
import { dispatchCreatePayment } from "../controller/payments"
import router from "../routers"
import { createSpan } from "../helpers/tracer"

export default {
  name: "CreateBalanceForm",
  components: {
    FormBalance: FormPayment
  },
  methods: {
    async createPayment (payment) {
      await createSpan("register-payment", async (span) => {
        await dispatchCreatePayment(span, payment)
        router.push({ name: "view-payments" })
      })
    }
  }
}
</script>
