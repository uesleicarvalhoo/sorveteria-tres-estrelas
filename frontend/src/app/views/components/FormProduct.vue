<template>
  <main-section>
    <title-sub-bar :icon="mdiBallotOutline" :title="title" />
    <card-component title="Registro" :icon="mdiBallot" @submit.prevent="submit" form>
      <field label="Nome">
        <control placeholder="Picole de {sabor}" type="text" v-model="form.name" required />
      </field>

      <field label="Preço de Varejo">
        <control type="number" placeholder="00.00" min="0" step="0.05" v-model="form.price_varejo" required />
      </field>

      <field label="Preço de Atacado">
        <control type="number" placeholder="00.00" min="0" step="0.05" v-model="form.price_atacado" required />
      </field>
      <field label="Qtd. para venda em atacado">
        <control type="number" placeholder="0" v-model="form.atacado_amount" min="0" step="1" required />
      </field>

      <divider />

      <jb-buttons>
        <jb-button type="submit" color="info" label="Confirmar" />
        <jb-button type="reset" color="info" outline label="Limpar" />
      </jb-buttons>
    </card-component>
  </main-section>
</template>

<script>
import { reactive, ref } from "vue"
import {
  mdiBallot,
  mdiBallotOutline,
  mdiAccount,
  mdiMail,
  mdiCheck
} from "@mdi/js"
import MainSection from "./MainSection.vue"
import CardComponent from "./CardComponent.vue"
import Divider from "./Divider.vue"
import JbButton from "./JbButton.vue"
import JbButtons from "./JbButtons.vue"
import Field from "./Field.vue"
import Control from "./Control.vue"
import TitleSubBar from "./TitleSubBar.vue"

export default {
  name: "FormProduct",
  components: {
    TitleSubBar,
    Divider,
    MainSection,
    CardComponent,
    Field,
    Control,
    JbButton,
    JbButtons
  },
  props: {
    data: {
      type: Object,
      default: () =>
        reactive({
          id: null,
          name: null,
          price_varejo: null,
          price_atacado: null,
          atacado_amount: null,
        })
    },
    title: { type: String, default: () => "Formulário de produtos" },
  },
  emits: ["submit"],

  setup(props, { emit }) {
    const warningModal = reactive({
      active: false,
      text: ""
    })

    const form = ref(props.data)

    const submit = () => {
      const item = {
        id: form.value.id,
        name: form.value.name,
        atacado_amount: form.value.atacado_amount,
        price_varejo: form.value.price_varejo,
        price_atacado: form.value.price_atacado
      }
      emit("submit", item)

      if (form.value.id === null) {
        reset()
      }
    }

    const reset = () => {
      form.value.name = null
      form.value.price_varejo = null
      form.value.price_atacado = null
      form.value.atacado_amount = null
    }

    return {
      form,
      submit,
      warningModal,
      mdiBallot,
      mdiBallotOutline,
      mdiAccount,
      mdiMail,
      mdiCheck
    }
  }
}
</script>
