import { productsService } from "../../services"
import { context } from "../helpers/context"
import { dispatchNotification, dispatchApiError } from "./notification"

/* Products */
export const dispatchCreateProduct = async (payload) => {
  try {
    await productsService.create(context.state.accessToken, payload)
    await dispatchNotification("Produto cadastrado", "Produto cadastrado com sucesso!", "success")
    await dispatchGetProducts()
  } catch (error) {
    await dispatchApiError(error)
  }
}

export const dispatchGetProducts = async () => {
  try {
    const products = await productsService.getAll(context.state.accessToken)
    context.commit("products", products)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export const dispatchUpdateProduct = async (payload) => {
  try {
    await productsService.update(context.state.accessToken, payload)
    dispatchNotification("Atualização do produto", "Dados do produto atualizados com sucesso!", "success")
    await dispatchGetProducts()
  } catch (error) {
    await dispatchApiError(error)
  }
}

export const dispatchRemoveProduct = async (payload) => {
  try {
    await productsService.delete(context.state.accessToken, payload.id)
    dispatchNotification("Exclusão do produto", "Produto excluido com sucesso!", "success")
    await dispatchGetProducts()
  } catch (error) {
    await dispatchApiError(error)
  }
}
