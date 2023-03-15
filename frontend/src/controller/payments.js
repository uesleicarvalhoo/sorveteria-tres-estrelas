import { context } from "../helpers/context"
import { paymentsService } from "../services/payments"
import { dispatchGetSales } from "./sales"
import { dispatchNotification, dispatchApiError } from "./notification"

export async function dispatchGetPayments (span) {
  try {
    const payments = await paymentsService.getAll(span)
    context.commit("payments", payments)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchGetPaymentsBetween (span, start, end) {
  try {
    const payments = await paymentsService.getBetween(span, start, end)
    context.commit("payments", payments)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchRemovePayment (span, payload) {
  try {
    await paymentsService.deletePayment(span, payload.id)
    await dispatchGetPayments()
    await dispatchNotification("Remoção do pagamento", "Pagamento removido com sucesso!", "success")
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchUpdatePayment (span, payload) {
  try {
    await paymentsService.updatePayment(span, payload)
    await dispatchGetPayments(span)
    await dispatchNotification("Atualização do pagamento", "Pagamento atualizado com sucesso!", "success")
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchCreatePayment (span, payload) {
  try {
    await paymentsService.registerPayment(span, payload)
    await dispatchNotification("Operação registrada", "Operação registrada com sucesso!", "success")
    await dispatchGetSales(span)
  } catch (error) {
    await dispatchApiError(error)
  }
}
