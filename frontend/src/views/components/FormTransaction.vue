<template>
  <main-section>
    <title-sub-bar :icon="mdiBallotOutline" title="Novo registro de Caixa" />
    <card-component :title="title" :icon="mdiBallot" @submit.prevent="submit" form>
      <field label="Tipo de Pagamento">
        <control :options="transactionTypes" v-model="form.transactionType" required />
      </field>
      <field label="Valor">
        <control type="number" placeholder="00.00" v-model="form.value" step="0.01" required />
      </field>
      <divider />
      <field label="Descrição da movimentação">
        <control type="textarea" v-model="form.description" required />
      </field>
      <divider />
      <jb-buttons>
        <jb-button type="submit" color="info" label="Confirmar" />
        <jb-button type="reset" color="info" outline label="Limpar" />
      </jb-buttons>
    </card-component>
  </main-section>
  <modal-box v-model="warningModal.active" large-title="Ops!" button="warning" buttonLabel="Ok" shake>
    <p v-text="warningModal.text"></p>
  </modal-box>
</template>

<script>
import { reactive, ref, computed } from "vue"
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
import ModalBox from "./ModalBox.vue"
import { context } from "../../helpers"

export default {
  name: "FormPayment",
  components: {
    TitleSubBar,
    Divider,
    MainSection,
    CardComponent,
    ModalBox,
    Field,
    Control,
    JbButton,
    JbButtons
  },
  props: {
    title: { type: String, default: "Nova Transação" },
    data: {
      type: Object,
      default: () =>
        reactive({
          id: null,
          value: null,
          transactionType: null,
          description: null
        })
    }
  },
  emits: ["submit"],

  setup (props, { emit }) {
    const warningModal = reactive({
      active: false,
      text: ""
    })

    const transactionTypes = computed(() => context.state.transactionTypes)

    const form = ref(props.data)

    const submit = () => {
      const data = {
        id: form.value.id,
        value: form.value.value,
        type: form.value.transactionType,
        description: form.value.description
      }
      emit("submit", data)

      if (form.value.id === null) {
        reset()
      }
    }

    const reset = () => {
      form.value.transactionType = null
      form.value.value = null
      form.value.description = ""
    }

    return {
      form,
      submit,
      transactionTypes,
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
