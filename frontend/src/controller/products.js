import { productsService } from "../services"
import { context } from "../helpers/context"
import { dispatchNotification, dispatchApiError } from "./notification"

/* Products */
export const dispatchCreateProduct = async (span, payload) => {
  try {
    await productsService.create(span, payload)
    await dispatchNotification("Produto cadastrado", "Produto cadastrado com sucesso!", "success")
    await dispatchGetProducts(span)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export const dispatchGetProducts = async (span) => {
  try {
    const products = await productsService.getAll(span)
    context.commit("products", products)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export const dispatchUpdateProduct = async (span, payload) => {
  try {
    await productsService.update(span, payload)
    dispatchNotification("Atualização do produto", "Dados do produto atualizados com sucesso!", "success")
    await dispatchGetProducts(span)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export const dispatchRemoveProduct = async (span, payload) => {
  try {
    await productsService.delete(span, payload.id)
    dispatchNotification("Exclusão do produto", "Produto excluido com sucesso!", "success")
    await dispatchGetProducts(span)
  } catch (error) {
    await dispatchApiError(error)
  }
}
