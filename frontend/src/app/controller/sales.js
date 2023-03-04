import { context } from "../helpers/context"
import { salesService } from "../../services"
import { dispatchNotification, dispatchApiError } from "./notification"

export async function dispatchGetSales () {
  try {
    const sales = await salesService.getAll(context.state.accessToken)
    context.commit("sales", sales)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchRegisterSale (payload) {
  try {
    await salesService.registerSale(context.state.accessToken, payload)
    await dispatchNotification("Venda registrada", "Venda registrada com sucesso!", "success")
    await dispatchGetSales()
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchRemoveSale (sale) {
  try {
    await salesService.delete(context.state.accessToken, sale.id)
    await dispatchNotification("Exclusão da venda", "Venda excluida com sucesso!", "success")
    await dispatchGetSales()
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchUpdateSale (sale) {
  try {
    await salesService.update(context.state.accessToken, sale)
    await dispatchGetSales
    await dispatchNotification("Atualização da venda", "Venda atualizada com sucesso!", "success")
  } catch (error) {
    await dispatchApiError(error)
  }
}
