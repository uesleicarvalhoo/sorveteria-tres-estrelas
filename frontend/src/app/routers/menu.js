import {
  mdiCart,
  mdiHome,
  mdiCurrencyUsd,
  mdiTableOfContents
} from "@mdi/js"

export default [
  "",
  [
    {
      to: { name: "home" },
      icon: mdiHome,
      label: "Dashboard"
    },
    {
      label: "Produtos",
      icon: mdiTableOfContents,
      menu: [
        {
          label: "Estoque",
          to: { name: "view-products" }
        },
        {
          label: "Novo produto",
          to: { name: "create-product" }
        }
      ]
    },
    {
      label: "Vendas",
      icon: mdiCart,
      menu: [
        {
          label: "Visualizar",
          to: { name: "view-sales" }
        },
        {
          label: "Registrar",
          to: { name: "create-sale" }
        }
      ]
    },
    {
      label: "Pagamentos",
      icon: mdiCurrencyUsd,
      menu: [
        {
          label: "Visualizar",
          to: { name: "view-payments" }
        },
        {
          label: "Registrar",
          to: { name: "register-payment" }
        }
      ]
    }
  ]
]
