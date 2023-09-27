<template>
  <hero-bar>Transações</hero-bar>
  <main-section>
    <card-component class="mb-6" has-table>
      <table-transaction v-on:remove="removeTransaction" v-on:view="view" />
    </card-component>
  </main-section>
</template>

<script>
import { mdiMonitorCellphone, mdiTableBorder } from "@mdi/js"
import MainSection from "./components/MainSection.vue"
import TableTransaction from "./components/TableTransaction.vue"
import CardComponent from "./components/CardComponent.vue"
import HeroBar from "./components/HeroBar.vue"
import { dispatchGetTransactions, dispatchRemoveTransaction } from "../controller/transaction"
import { reactive } from "vue"
import { createSpan } from "../helpers/tracer"

export default {
  name: "ViewTransactions",
  components: {
    MainSection,
    HeroBar,
    CardComponent,
    TableTransaction
  },
  methods: {
    async view (transaction) {
      this.modal.active = true
      Object.assign(this.modal.data, transaction)
    },
    async removeTransaction (transaction) {
      await createSpan("delete-transaction", async (span) => {
        await dispatchRemoveTransaction(span, transaction)
      })
    }
  },
  async created () {
    await createSpan("view-transaction", async (span) => {
      await dispatchGetTransactions(span)
    })
  },
  setup () {
    const modal = reactive({
      active: false,
      data: {
        id: null,
        created_at: "",
        description: "",
        value: 0
      }
    })

    return {
      modal,
      mdiMonitorCellphone,
      mdiTableBorder
    }
  }
}
</script>
