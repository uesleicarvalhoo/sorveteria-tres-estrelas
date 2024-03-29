export const itemsPerPage = process.env.VUE_APP_ITEMS_PER_PAGE || 10
export const apiUrl = process.env.VUE_APP_API_URL || `${window.location.protocol}//${window.location.host}/api`
export const serviceName = process.env.VUE_APP_SERVICE_NAME || "frontend"
export const serviceVersion = process.env.VUE_APP_SERVICE_VERSION || "0.0.0"
export const otelEndpoint = process.env.VUE_APP_OTEL_URL || "http://localhost:9411/api/v2/spans"
