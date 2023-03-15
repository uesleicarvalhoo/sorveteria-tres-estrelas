<template>
  <hero-bar>Pagamentos</hero-bar>
  <main-section>
    <card-component class="mb-6" has-table>
      <table-payment v-on:remove="removePayment" v-on:view="view" />
    </card-component>
  </main-section>
  <modal-view v-model="modal.active">
    <form-payment title="Atualizar pagamento" :data="modal.data" v-on:submit="updatePayment" />
  </modal-view>
</template>

<script>
import { mdiMonitorCellphone, mdiTableBorder } from "@mdi/js"
import MainSection from "./components/MainSection.vue"
import TablePayment from "./components/TablePayment.vue"
import CardComponent from "./components/CardComponent.vue"
import HeroBar from "./components/HeroBar.vue"
import { dispatchGetPayments, dispatchRemovePayment, dispatchUpdatePayment } from "../controller/payments"
import { reactive } from "vue"
import FormPayment from "./components/FormPayment.vue"
import ModalView from "./components/ModalView.vue"
import { createSpan } from "../helpers/tracer"

export default {
  name: "ViewPayments",
  components: {
    MainSection,
    HeroBar,
    CardComponent,
    TablePayment,
    ModalView,
    FormPayment
  },
  methods: {
    async view (payment) {
      this.modal.active = true
      Object.assign(this.modal.data, payment)
    },
    async updatePayment (payment) {
      await createSpan("update-payment", async (span) => {
        await dispatchUpdatePayment(span, payment)
      })
    },
    async removePayment (payment) {
      await createSpan("delete-payment", async (span) => {
        await dispatchRemovePayment(span, payment)
      })
    }
  },
  async created () {
    await createSpan("view-payments", async (span) => {
      await dispatchGetPayments(span)
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
