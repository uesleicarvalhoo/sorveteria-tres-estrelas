<template>
  <hero-bar>Estoque de Produtos</hero-bar>
  <main-section>
    <card-component class="mb-6" has-table>
      <table-products v-on:view="viewProduct" v-on:remove="removeProduct" />
    </card-component>
  </main-section>
  <modal-view v-model="modal.active">
    <form-product title="Atualizar produto" :data="modal.data" v-on:submit="updateProduct" />
  </modal-view>
</template>

<script>
import { reactive } from "vue"
import { mdiMonitorCellphone, mdiTableBorder } from "@mdi/js"
import MainSection from "./components/MainSection.vue"
import TableProducts from "./components/TableProduct.vue"
import CardComponent from "./components/CardComponent.vue"
import HeroBar from "./components/HeroBar.vue"
import FormProduct from "./components/FormProduct.vue"
import ModalView from "./components/ModalView.vue"
import { dispatchGetProducts, dispatchRemoveProduct, dispatchUpdateProduct } from "../controller/products"
import { createSpan } from "../helpers/tracer"

export default {
  name: "ViewItem",
  components: {
    MainSection,
    HeroBar,
    CardComponent,
    TableProducts,
    FormProduct,
    ModalView
  },
  methods: {
    viewProduct (item) {
      this.modal.active = true
      Object.assign(this.modal.data, item)
    },
    async removeProduct (item) {
      await createSpan("delete-product", async (span) => {
        await dispatchRemoveProduct(span, item)
      })
    },
    async updateProduct (item) {
      await createSpan("update-product", async (span) => {
        await dispatchUpdateProduct(span, item)
      })
    }
  },
  async created () {
    await createSpan("view-products", async (span) => {
      await dispatchGetProducts(span)
    })
  },
  setup () {
    const modal = reactive({
      active: false,
      data: {
        id: null,
        name: "",
        price_varejo: 0,
        price_atacado: 0,
        atacado_amount: 0
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
