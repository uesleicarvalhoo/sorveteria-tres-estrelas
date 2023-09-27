<template>
  <form-transaction v-on:submit="createTransaction"></form-transaction>
</template>

<script>
import FormTransaction from "./components/FormTransaction.vue"
import { dispatchCreateTransaction } from "../controller/transaction" // TODO
import router from "../routers"
import { createSpan } from "../helpers/tracer"

export default {
  name: "RegisterTransactionForm",
  components: {
    FormTransaction: FormTransaction
  },
  methods: {
    async createTransaction (balance) {
      await createSpan("register-transaction", async (span) => {
        await dispatchCreateTransaction(span, balance)
        router.push({ name: "view-transaction" })
      })
    }
  }
}
</script>
