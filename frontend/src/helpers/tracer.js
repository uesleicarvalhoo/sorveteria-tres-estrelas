import { BatchSpanProcessor } from "@opentelemetry/sdk-trace-base"
import { WebTracerProvider } from "@opentelemetry/sdk-trace-web"
import { ZoneContextManager } from "@opentelemetry/context-zone"
import { SpanStatusCode, trace } from "@opentelemetry/api"
import { Resource } from "@opentelemetry/resources"
import { SemanticResourceAttributes } from "@opentelemetry/semantic-conventions"
import { ZipkinExporter } from "@opentelemetry/exporter-zipkin"
import { serviceName, serviceVersion, otelEndpoint } from "../config/index"

const provider = new WebTracerProvider({
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: serviceName,
    [SemanticResourceAttributes.SERVICE_VERSION]: serviceVersion
  })
})

const zipkinExporter = new ZipkinExporter({
  serviceName: serviceName,
  url: otelEndpoint
})

provider.addSpanProcessor(new BatchSpanProcessor(zipkinExporter))
provider.register({
  contextManager: new ZoneContextManager()

})

export const tracer = trace.getTracer(serviceName)

export const createSpan = async (name, fn) => {
  tracer.startActiveSpan(name, async (span) => {
    try {
      return await fn(span)
    } catch (err) {
      span.setStatus({
        code: SpanStatusCode.ERROR,
        message: err.message
      })
      throw err
    } finally {
      span.end()
    }
  })
}
