<template>
    <Datepicker v-model="date" :format="format" range @cleared="emit('clear')" />
</template>

<script>
import { ref } from "vue"
import Datepicker from "@vuepic/vue-datepicker"
import "@vuepic/vue-datepicker/dist/main.css"

export default {
  components: {
    Datepicker
  },
  props: {
    startDate: new Date(),
    endDate: new Date()
  },
  emits: ["submit", "clear"],
  setup (props, { emit }) {
    const date = ref(new Date())

    const format = ([start, end]) => {
      emit("submit", start, end)
      return `${start.toLocaleDateString("pt-br")} ~ ${end.toLocaleDateString("pt-br")}`
    }

    return {
      date,
      format,
      emit
    }
  }
}
</script>
