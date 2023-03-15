import { context } from "../helpers/context"
import { cashFlowService } from "../services/cashflow"
import { dispatchApiError } from "./notification"

export async function dispatchGetCashFlow (span) {
  try {
    const cashFlow = await cashFlowService.getAll(span)
    context.commit("cashFlow", cashFlow)
  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchGetCashFlowBetween (span, start, end) {
  try {
    const cashFlow = await cashFlowService.getBetween(span, start, end, context.state.accessToken)
    context.commit("cashFlow", cashFlow)
  } catch (error) {
    await dispatchApiError(error)
  }
}
