<template>
  <full-screen-section bg="login" v-slot="{ cardClass, cardRounded }">
    <card-component :class="cardClass" :rounded="cardRounded" @submit.prevent="submit" form>
      <field label="Login" help="Por favor, informe o seu email">
        <control v-model="form.email" :icon="mdiAccount" name="login" autocomplete="email" />
      </field>

      <field label="Password" help="Por favor, informe a sua senha">
        <control v-model="form.password" :icon="mdiAsterisk" type="password" name="password"
          autocomplete="current-password" />
      </field>

      <divider />

      <jb-buttons>
        <jb-button type="submit" color="info" label="Login" />
        <jb-button to="/" color="info" outline label="Back" />
      </jb-buttons>
    </card-component>
  </full-screen-section>
</template>

<script>
import { mdiAccount, mdiAsterisk } from "@mdi/js"
import FullScreenSection from "./components/FullScreenSection.vue"
import CardComponent from "./components/CardComponent.vue"
import Field from "./components/Field.vue"
import Control from "./components/Control.vue"
import Divider from "./components/Divider.vue"
import JbButton from "./components/JbButton.vue"
import JbButtons from "./components/JbButtons.vue"
import { dispatchLogin } from "../controller/auth"
import { reactive } from "@vue/reactivity"
import { createSpan } from "../helpers/tracer"

export default {
  name: "Login",
  components: {
    FullScreenSection,
    CardComponent,
    Field,
    Control,
    Divider,
    JbButton,
    JbButtons
  },
  props: {
    email: String,
    password: String
  },
  methods: {
    async submit () {
      await createSpan("login", async (span) => {
        await dispatchLogin(span, this.form.email, this.form.password)
      })
    }
  },
  setup () {
    const form = reactive({ email: null, password: null })

    return {
      form,
      mdiAccount,
      mdiAsterisk
    }
  }
}
</script>
