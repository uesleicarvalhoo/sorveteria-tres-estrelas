<template>
  <table>
    <thead>
      <tr>
        <th class="text-center">Código</th>
        <th class="text-center">Descrição</th>
        <th class="text-center">Preço Varejo</th>
        <th class="text-center">Preço Atacado</th>
        <th class="text-center">Quantidade Atacado</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in itemsPaginated" :key="item.id">
        <td data-label="Código">{{ item.code }}</td>
        <td class="text-center" data-label="Descrição">{{ item.name }}</td>
        <td class="text-center" data-label="Valor Varejo">
          R$ {{ item.price_varejo }}
        </td>
        <td class="text-center" data-label="Valor Atacado">
          R$ {{ item.price_atacado }}
        </td>
        <td class="text-center" data-label="Quantidade para Atacado">
          {{ item.atacado_amount }}
        </td>
        <td class="actions-cell">
          <jb-buttons type="justify-start lg:justify-end" no-wrap v-if="actions">
            <jb-button class="mr-3" color="success" :icon="mdiEye" small @click="emitEvent('view', item)" />
            <jb-button color="danger" :icon="mdiTrashCan" small @click="emitEvent('remove', item)" />
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
  name: "TableItem",
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

    const products = computed(() => context.state.products)

    const perPage = ref(itemsPerPage)

    const currentPage = ref(0)

    const itemsPaginated = computed(() =>
      products.value.slice(
        perPage.value * currentPage.value,
        perPage.value * (currentPage.value + 1)
      )
    )

    const numPages = computed(() =>
      Math.ceil(products.value.length / perPage.value)
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
