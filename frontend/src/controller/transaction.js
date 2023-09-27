import { context } from "../helpers/context"
import { transactionService } from "../services/transaction"
import { dispatchNotification, dispatchApiError } from "./notification"

export async function dispatchGetTransactions (span) {
  try {
    const transactions = await transactionService.getAll(span)
    context.commit("transactions", transactions)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchRemoveTransaction (span, payload) {
  try {
    await transactionService.deleteTransaction(span, payload.id)
    await dispatchNotification("Remoção da transação", "Transação removida com sucesso!", "success")
    await dispatchGetTransactions(span)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchCreateTransaction (span, payload) {
  try {
    await transactionService.registerTransaction(span, payload)
    await dispatchGetTransactions(span)
    await dispatchNotification("Transação registrada", "Transação registrada com sucesso!", "success")
  } catch (error) {
    await dispatchApiError(error)
  }
}
