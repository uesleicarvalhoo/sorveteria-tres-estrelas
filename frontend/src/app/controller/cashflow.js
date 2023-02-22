import { context } from "../helpers/context"
import { cashFlowService } from "../../services/cashflow"
import { dispatchApiError } from "./notification"

export async function dispatchGetCashFlow() {
  try {
    const cashFlow = await cashFlowService.getAll(context.state.accessToken)
    context.commit("cashFlow", cashFlow)

  } catch (error) {
    await dispatchApiError(error)
  }
}

export async function dispatchGetCashFlowBetween(start, end) {
  try {
    const cashFlow = await cashFlowService.getBetween(start, end, context.state.accessToken)
    context.commit("cashFlow", cashFlow)
  } catch (error) {
    await dispatchApiError(error)
  }
}
