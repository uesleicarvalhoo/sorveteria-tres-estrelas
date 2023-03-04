<template>
  <main-section>
    <title-sub-bar :icon="mdiBallotOutline" :title="title" />
    <card-component title="Registro" :icon="mdiBallot" @submit.prevent="submit" form>
      <field label="Tipo de Pagamento">
        <control :options="paymentTypes" v-model="form.payment_type" required />
      </field>

      <field label="Descrição">
        <control placeholder="Descrição da venda" type="textarea" v-model="form.description" required />
      </field>

      <divider />

      <field label="Produtos">
        <control :options="products" v-model="form.product" />
        <control v-model="form.amount" placeholder="Quantidade do produto" type="number" min="1" step="1" />
        <jb-button color="success" label="Adicionar" :icon="mdiCheck" small @click="addItem()" />
      </field>

      <table>
        <thead>
          <tr>
            <th class="text-center">Descrição</th>
            <th class="text-center">Quantidade</th>
            <th class="text-center">Preço unitário</th>
            <th class="text-center">Sub total</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in form.items">
            <td data-label="Descrição" class="text-center">{{ item.name }}</td>
            <td data-label="Código" class="text-center">{{ item.amount }}</td>
            <td data-label="Valor" class="text-center">R$ {{ item.unit_price.toFixed(2) }}</td>
            <td data-label="Quantidade" class="text-center">
              R$ {{ (item.amount * item.unit_price).toFixed(2) }}
            </td>
            <td class="actions-cell items-center justify-between">
              <jb-buttons no-wrap>
                <jb-button color="danger" :icon="mdiTrashCan" small @click="removeItem(index)" />
              </jb-buttons>
            </td>
          </tr>
        </tbody>
      </table>

      <divider />

      <jb-buttons>
        <jb-button color="info" type="submit" label="Confirmar" />
        <jb-button type="reset" color="info" outline label="Limpar" />
      </jb-buttons>
    </card-component>
  </main-section>

  <modal-box v-model="warningModal.active" large-title="Ops!" button="warning" buttonLabel="Ok" shake>
    <p v-text="warningModal.text"></p>
  </modal-box>
</template>

<script>
import { computed, reactive, ref } from "vue"
import { mdiBallot, mdiBallotOutline, mdiCheck, mdiTrashCan } from "@mdi/js"
import MainSection from "./MainSection.vue"
import CardComponent from "./CardComponent.vue"
import Divider from "./Divider.vue"
import JbButton from "./JbButton.vue"
import JbButtons from "./JbButtons.vue"
import Field from "./Field.vue"
import Control from "./Control.vue"
import TitleSubBar from "./TitleSubBar.vue"
import ModalBox from "./ModalBox.vue"
import { dispatchGetProducts } from "../../controller/products"
import { useStore } from "vuex"

export default {
  name: "FormSale",
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
    title: { type: String, default: () => "Formulário de vendas" },
    data: {
      type: Object,
      default: () =>
        reactive({
          id: null,
          payment_type: null,
          description: null,
          product: null,
          amount: 1,
          items: []
        })
    }
  },
  emits: ["submit"],
  async created() {
    await dispatchGetProducts()
  },
  setup(props, { emit }) {
    const form = ref(props.data)

    const context = useStore()
    const products = computed(() => context.state.products)

    const paymentTypes = computed(() => context.state.paymentTypes)

    const warningModal = reactive({
      active: false,
      text: ""
    })

    const getItemFromForm = (name) => {
      const filtredProducts = form.value.items.filter((el) => el.name === name)
      if (filtredProducts.length > 0) {
        return filtredProducts[0]
      } else {
        return null
      }
    }

    const getProductPrice = (product, amount) => {
      if (amount >= product.atacado_amount) {
        return product.price_atacado
      } else {
        return product.price_varejo
      }
    }

    const addItem = () => {
      if (!form.value.amount | (form.value < 1)) {
        warningModal.active = true
        warningModal.text = "A quantidade do produto precisa ser maior que 0!"
        return
      }

      if (form.value.product.id) {
        const product = Object.assign(form.value.product)
        const item = getItemFromForm(product.name)

        if (item !== null) {
          item.amount = item.amount + Number(form.value.amount)
          item.unit_price = getProductPrice(product, item.amount)
        } else {
          form.value.items.push({
            id: product.id,
            name: product.name,
            amount: Number(form.value.amount),
            unit_price: getProductPrice(product, Number(form.value.amount))
          })
        }
      }
    }

    const removeItem = (index) => {
      form.value.items.splice(index, 1)
    }

    const submit = () => {
      if (form.value.items.length === 0) {
        warningModal.active = true
        warningModal.text = "Adicione pelo menos um item ao carrinho!"
      } else {
        const details = []

        form.value.items.forEach((el) => {
          details.push({
            item_id: el.id,
            amount: el.amount
          })
        })

        const data = {
          id: form.value.id,
          payment_type: form.value.payment_type,
          description: form.value.description,
          items: details
        }
        emit("submit", data)
        if (data.id === null) {
          reset()
        }
      }
    }

    const reset = () => {
      form.value.payment_type = null
      form.value.description = null
      form.value.amount = null
      form.value.product = null
      form.value.items = []
    }

    return {
      form,
      submit,
      products,
      addItem,
      removeItem,
      warningModal,
      paymentTypes,
      mdiBallot,
      mdiBallotOutline,
      mdiCheck,
      mdiTrashCan
    }
  }
}
</script>
