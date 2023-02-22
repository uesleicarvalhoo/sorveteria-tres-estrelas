<template>
  <main-section>
    <div class="grid grid-cols-1 gap-6 lg:grid-cols-3 mb-6">
      <card-widget color="text-blue-500" :icon="mdiCartOutline" prefix="R$" :number="totalSales"
        label="Total de vendas" />
      <card-widget color="text-blue-500" :icon="mdiCurrencyUsd" prefix="R$" :number="totalPayments"
        label="Total de Pagamentos" />
      <card-widget color="text-red-500" :icon="mdiChartTimelineVariant" :number="balance" prefix="R$"
        label="Balanço geral" />
    </div>
    <calendar v-on:submit="getCashFlowBetween" v-on:clear="dispatchGetCashFlow"> </calendar>
    <card-component title="Movimentações do mês" has-table>
      <table-balance :actions="false" />
    </card-component>
  </main-section>
</template>

<script>
import { computed } from "vue"
import { useStore } from "vuex"
import {
  mdiCartOutline,
  mdiChartTimelineVariant,
  mdiCurrencyUsd
} from "@mdi/js"
import MainSection from "./components/MainSection.vue"
import CardWidget from "./components/CardWidget.vue"
import CardComponent from "./components/CardComponent.vue"
import TableBalance from "./components/TableBalance.vue"
import Calendar from "./components/Calendar.vue"
import { dispatchGetMe } from "../controller/users"
import { dispatchGetSales } from "../controller/sales"
import { dispatchGetCashFlow, dispatchGetCashFlowBetween } from "../controller/cashflow"

export default {
  name: "Home",
  components: {
    MainSection,
    CardComponent,
    TableBalance,
    CardWidget,
    Calendar
  },
  methods: {
    async getCashFlowBetween (start, end) {
      await dispatchGetCashFlowBetween(start, end)
    }
  },
  async created () {
    await dispatchGetMe()
    await dispatchGetSales()
    await dispatchGetCashFlow()
  },
  setup () {
    const context = useStore()

    const totalSales = computed(() => (context.state.cashFlow.total_sales))

    const totalPayments = computed(() => (context.state.cashFlow.total_sales))

    const balance = computed(() => context.state.cashFlow.balance)

    const darkMode = computed(() => context.state.darkMode)
    return {
      totalPayments,
      totalSales,
      balance,
      darkMode,
      mdiCartOutline,
      mdiCurrencyUsd,
      mdiChartTimelineVariant,
      dispatchGetCashFlow
    }
  }
}
</script>
