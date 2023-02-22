import { context } from "../helpers/context"
import { paymentsService } from "../../services/payments"
import { dispatchGetSales } from "./sales"
import { dispatchNotification } from "./notification"

export async function dispatchGetPayments() {
  try {
    const payments = await paymentsService.getAll(context.state.accessToken)
    context.commit("payments", payments)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchGetPaymentsBetween(start, end) {
  try {
    const payments = await paymentsService.getBetween(start, end, context.state.accessToken)
    context.commit("payments", payments)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchRemovePayment(payload) {
  try {
    await paymentsService.deletePayment(context.state.accessToken, payload.id)
    await dispatchGetPayments()
    await dispatchNotification("Remoção do pagamento", "Pagamento removido com sucesso!", "success")
  } catch {
    await dispatchApiError(error)
  }
}

export async function dispatchUpdatePayment(payload) {
  try {
    await paymentsService.updatePayment(context.state.accessToken, payload)
    await dispatchGetPayments()
    await dispatchNotification("Atualização do pagamento", "Pagamento atualizado com sucesso!", "success")
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchCreatePayment(payload) {
  try {
    await paymentsService.registerPayment(context.state.accessToken, payload)
    await dispatchNotification("Operação registrada", "Operação registrada com sucesso!", "success")
    await dispatchGetSales()
  } catch (error) {
    await dispatchApiError(error)
  }
}
