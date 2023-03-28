# Build stage
FROM golang:1.20-alpine AS builder

# Install packages
RUN apk update && \
    apk add --no-cache musl-dev gcc build-base libc-dev curl git tzdata ca-certificates && \
    update-ca-certificates && \
    echo "America/Sao_Paulo" > /etc/timezone

# Install dependencies
ENV GOPATH="$HOME/go"

WORKDIR $GOPATH/src
ADD go.mod go.sum $GOPATH/src/

RUN GOOS=linux go mod download

# Build backend

COPY . $GOPATH/src

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/backend main.go

# Deploy
FROM scratch

# Copy configurations
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=builder /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone

# Copy migrations and execulatable
COPY --from=builder /go/src/database/migrations/ /app/database/migrations/

WORKDIR /app
COPY --from=builder /go/bin/backend .

ENV HTTP_SERVER_PORT 80

EXPOSE ${HTTP_SERVER_PORT}

ENTRYPOINT ["/app/backend"]