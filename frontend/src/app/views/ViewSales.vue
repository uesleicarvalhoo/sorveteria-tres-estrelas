<template>
  <hero-bar>Vendas</hero-bar>
  <main-section>
    <card-component class="mb-6" has-table>
      <table-sale v-on:view="viewSale" v-on:remove="removeSale" />
    </card-component>
  </main-section>
  <modal-view v-model="modal.active">
    <form-sale title="Atualizar" v-on:submit="updateSale" :data="modal.data" />
  </modal-view>
</template>

<script>
import { reactive } from "vue"
import { mdiMonitorCellphone, mdiTableBorder } from "@mdi/js"
import MainSection from "./components/MainSection.vue"
import CardComponent from "./components/CardComponent.vue"
import HeroBar from "./components/HeroBar.vue"
import TableSale from "./components/TableSale.vue"
import FormSale from "./components/FormSale.vue"
import ModalView from "./components/ModalView.vue"
import {
  dispatchGetSales,
  dispatchRemoveSale
} from "../controller/sales"

import { dispatchNotification } from "../controller/notification"

export default {
  name: "ViewSale",
  components: {
    MainSection,
    HeroBar,
    CardComponent,
    TableSale,
    FormSale,
    ModalView
  },
  methods: {
    async removeSale(sale) {
      await dispatchRemoveSale(sale)
    },

    async updateSale(sale) {
      dispatchNotification(
        "Função desabilitada",
        "Ops! Ainda não configurei a função de atualizar as vendas, mas você pode remover e cadastrar novamente.",
        "warning")
    },

    viewSale(sale) {
      Object.assign(this.modal.data, sale)
      this.modal.active = true
    }
  },
  async created() {
    await dispatchGetSales()
  },
  setup() {
    const modal = reactive({
      active: false,
      data: {
        id: null,
        description: "",
        payment_type: "",
        total: 0,
        items: []
      }
    })

    return {
      mdiMonitorCellphone,
      mdiTableBorder,
      modal
    }
  }
}
</script>
