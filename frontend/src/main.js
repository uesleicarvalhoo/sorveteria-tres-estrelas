import { createApp } from "vue"

import App from "./app/App.vue"
import router from "./app/routers"
import { context } from "./app/helpers/context"

import "./app/styles/css/main.css"

/* Collapse mobile aside menu on route change */
router.beforeEach(to => {
  context.dispatch("asideMobileToggle", false)
  context.dispatch("asideLgToggle", false)
})

/* Default title tag */
const defaultDocumentTitle = "Sorveteria Três Estrelas"

router.afterEach(to => {
  /* Set document title from route meta */
  if (to.meta && to.meta.title) {
    document.title = `${defaultDocumentTitle} - ${to.meta.title}`
  } else {
    document.title = defaultDocumentTitle
  }

  /* Full screen mode */
  context.dispatch("fullScreenToggle", !!to.meta.fullScreen)
})

createApp(App).use(context).use(router).mount("#app")
