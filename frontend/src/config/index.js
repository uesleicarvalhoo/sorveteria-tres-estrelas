export const itemsPerPage = process.env.ITEMS_PER_PAGE || 10

const envApiUrl = process.env.VUE_APP_API_URL || `${window.location.protocol}//${window.location.host}/api`

export const apiUrl = envApiUrl

export const serviceName = process.env.SERVICE_NAME || "frontend"
export const serviceVersion = process.env.SERVICE_VERSION || "0.0.0"
export const otelEndpoint = process.env.OTEL_URL || "http://localhost:9411/api/v2/spans"
