import { createStore } from "vuex"

export const context = createStore({
  state: {
    /* User */
    user: {
      id: null,
      name: null,
      email: null
    },

    userAvatar: "https://avatars.dicebear.com/api/avataaars/example.svg?options[top][]=shortHair&options[accessoriesChance]=93",
    accessToken: null,
    refreshToken: null,
    expireTokenTime: null,

    /* Balance */
    cashFlow: {
      balance: 0,
      details: [],
      total_payments: 0,
      total_sales: 0
    },

    /* Products */
    products: [],

    /* Sale */
    sales: [],

    /* Transactions */
    transactions: [],

    /* Constants */
    transactionTypes: [
      "Venda",
      "Pagamento"
    ],

    /* Notification */
    showNotification: false,
    notification: {},

    /* Access info */
    loggedIn: false,

    /* fullScreen - fullscreen form layout (e.g. login page) */
    isFullScreen: false,

    /* Aside */
    isAsideMobileExpanded: false,
    isAsideLgActive: false,

    /* Dark mode */
    darkMode: false,

    /* Field focus with ctrl+k (to register only once) */
    isFieldFocusRegistered: false
  },
  mutations: {
    /* A fit-them-all commit */
    basic (state, payload) {
      state[payload.key] = payload.value
    },

    /* User */
    user (state, payload) {
      state.user = payload
    },
    loggedIn (state, paylolad) {
      state.loggedIn = paylolad
    },
    accessToken (state, payload) {
      state.accessToken = payload.token
      state.refreshToken = payload.token
      state.expireTokenTime = payload.expiration
    },
    sales (state, payload) {
      state.sales = payload
    },
    payments (state, payload) {
      state.payments = payload
    },
    cashFlow (state, payload) {
      state.cashFlow = payload
    },
    products (state, payload) {
      state.products = payload
    },
    notification (state, payload) {
      state.notification = payload
    },
    showNotification (state, payload) {
      state.showNotification = payload
    },
    transactions (state, payload) {
      state.transactions = payload
    }
  },
  actions: {
    asideMobileToggle ({ commit, state }, payload = null) {
      const isShow = payload !== null ? payload : !state.isAsideMobileExpanded

      document.getElementById("app").classList[isShow ? "add" : "remove"]("ml-60")

      document.documentElement.classList[isShow ? "add" : "remove"]("m-clipped")

      commit("basic", {
        key: "isAsideMobileExpanded",
        value: isShow
      })
    },

    asideLgToggle ({ commit, state }, payload = null) {
      commit("basic", { key: "isAsideLgActive", value: payload !== null ? payload : !state.isAsideLgActive })
    },

    fullScreenToggle ({ commit, state }, value) {
      commit("basic", { key: "isFullScreen", value })

      document.documentElement.classList[value ? "add" : "remove"]("full-screen")
    },

    darkMode ({ commit, state }) {
      const value = !state.darkMode

      document.documentElement.classList[value ? "add" : "remove"]("dark")

      commit("basic", {
        key: "darkMode",
        value
      })
    }
  },
  modules: {
  }
})
