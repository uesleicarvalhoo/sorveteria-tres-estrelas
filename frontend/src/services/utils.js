import { context } from "../helpers/context"

export function getContextHeaders (span) {
  const spanContext = span.spanContext()
  const headers = {
    traceparent: `00-${spanContext.traceId}-${spanContext.spanId}-01`
  }

  if (context.state.loggedIn) {
    headers.Authorization = `Bearer ${context.state.accessToken}`
  }

  return headers
}
