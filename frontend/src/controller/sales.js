import { context } from "../helpers/context"
import { salesService } from "../services"
import { dispatchNotification, dispatchApiError } from "./notification"

export async function dispatchGetSales (span) {
  try {
    const sales = await salesService.getAll(span)
    context.commit("sales", sales)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchRegisterSale (span, payload) {
  try {
    await salesService.registerSale(span, payload)
    await dispatchNotification("Venda registrada", "Venda registrada com sucesso!", "success")
    await dispatchGetSales(span)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchRemoveSale (span, sale) {
  try {
    await salesService.delete(span, sale.id)
    await dispatchNotification("Exclusão da venda", "Venda excluida com sucesso!", "success")
    await dispatchGetSales(span)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchUpdateSale (span, sale) {
  try {
    await salesService.update(span, sale)
    await dispatchGetSales(span)
    await dispatchNotification("Atualização da venda", "Venda atualizada com sucesso!", "success")
  } catch (error) {
    await dispatchApiError(error)
  }
}
