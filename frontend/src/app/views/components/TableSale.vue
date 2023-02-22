<template>
  <table>
    <thead>
      <tr>
        <th class="text-center">Descrição</th>
        <th class="text-center">Tipo de pagamento</th>
        <th class="text-center">Total</th>
        <th class="text-center">Data</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="sale in itemsPaginated" :key="sale.id">
        <td class="text-center" data-label="Descrição">
          {{ sale.description }}
        </td>
        <td class="text-center" data-label="Tipo de pagamento">
          {{ sale.payment_type }}
        </td>
        <td class="text-center" data-label="Total">
          R$
          {{ sale.total.toFixed(2) }}
        </td>
        <td class="text-center" data-label="Data">
          {{ new Date(sale.date).toLocaleDateString() }}
        </td>
        <td class="actions-cell" v-if="actions">
          <jb-buttons type="justify-start lg:justify-end" no-wrap>
            <jb-button class="mr-3" color="success" :icon="mdiEye" :to="{ name: 'view-sales' }" small
              @click="emitEvent('view', sale)" />
            <jb-button color="danger" :icon="mdiTrashCan" small @click="emitEvent('remove', sale)" />
          </jb-buttons>
        </td>
      </tr>
    </tbody>
  </table>
  <div class="table-pagination">
    <level>
      <jb-buttons>
        <jb-button v-for="page in pagesList" @click="currentPage = page" :active="page === currentPage"
          :label="page + 1" :key="page" :outline="darkMode" small />
      </jb-buttons>
      <small>Pagina {{ currentPageHuman }} de {{ numPages }}</small>
    </level>
  </div>
</template>

<script>
import { computed, ref } from "vue"
import { useStore } from "vuex"
import { mdiEye, mdiTrashCan } from "@mdi/js"
import Level from "./Level.vue"
import JbButtons from "./JbButtons.vue"
import JbButton from "./JbButton.vue"
import { itemsPerPage } from "../../../config"

export default {
  name: "TableSale",
  components: {
    Level,
    JbButtons,
    JbButton
  },
  props: {
    actions: { type: Boolean, default: true }
  },
  emits: ["view", "remove"],

  setup (props, { emit }) {
    const context = useStore()

    const darkMode = computed(() => context.state.darkMode)

    const items = computed(() => context.state.sales)

    const perPage = ref(itemsPerPage)

    const currentPage = ref(0)

    const itemsPaginated = computed(() =>
      items.value.slice(
        perPage.value * currentPage.value,
        perPage.value * (currentPage.value + 1)
      )
    )

    const numPages = computed(() =>
      Math.ceil(items.value.length / perPage.value)
    )

    const currentPageHuman = computed(() => currentPage.value + 1)

    const pagesList = computed(() => {
      const pagesList = []

      for (let i = 0; i < numPages.value; i++) {
        pagesList.push(i)
      }

      return pagesList
    })

    const emitEvent = (event, data) => {
      emit(event, data)
    }

    return {
      darkMode,
      emitEvent,
      currentPage,
      currentPageHuman,
      numPages,
      itemsPaginated,
      pagesList,
      mdiEye,
      mdiTrashCan
    }
  }
}
</script>
