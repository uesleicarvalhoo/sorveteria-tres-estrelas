import { createRouter, createWebHistory } from "vue-router"
import Home from "../views/Home"
import { loginRequired, refreshToken } from "./middlewares/auth"

const routes = [
  {
    // Document title tag
    // We combine it with defaultDocumentTitle set in `src/main.js` on router.afterEach hook
    meta: {
      title: "Dashboard"
    },
    path: "/",
    name: "home",
    component: Home
  },
  {
    meta: {
      title: "Produtos"
    },
    path: "/produtos/cadastro",
    name: "create-product",
    component: () => import("../views/CreateProduct")
  },
  {
    meta: {
      title: "Vendas"
    },
    path: "/vendas/cadastro",
    name: "create-sale",
    component: () => import("../views/RegisterSale")
  },
  {
    meta: {
      title: "Vendas"
    },
    path: "/vendas/",
    name: "view-sales",
    component: () => import("../views/ViewSales")
  },
  {
    meta: {
      title: "Estoque"
    },
    path: "/produtos/",
    name: "view-products",
    component: () => import("../views/ViewProduct")
  },
  {
    meta: {
      title: "Pagamentos"
    },
    path: "/pagamentso/registrar",
    name: "register-payment",
    component: () => import("../views/RegisterPayment")
  },
  {
    meta: {
      title: "Pagamentos"
    },
    path: "/pagamentos",
    name: "view-payments",
    component: () => import("../views/ViewPayments")
  },
  {
    meta: {
      title: "Login",
      fullScreen: true
    },
    path: "/login",
    name: "login",
    component: () => import("../views/Login")
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior (to, from, savedPosition) {
    return savedPosition || { top: 0 }
  }
})

router.beforeEach(refreshToken)
router.beforeEach(loginRequired)

export default router
export { router }
