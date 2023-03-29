FROM node:lts-alpine as builder

WORKDIR /app

# Copy the package.json and install dependencies
COPY package.json ./

COPY yarn.lock ./

RUN yarn install --frozen-lockfile

# Copy rest of the files
COPY . .

# Build the project
ARG VUE_APP_ITEMS_PER_PAGE=10
ARG VUE_APP_API_URL=""
ARG VUE_APP_SERVICE_NAME="frontend"
ARG VUE_APP_SERVICE_VERSION="0.0.0"
ARG VUE_APP_OTEL_URL="http://localhost:9411/api/v2/spans"

ENV VUE_APP_ITEMS_PER_PAGE=${VUE_APP_ITEMS_PER_PAGE}
ENV VUE_APP_API_URL=${VUE_APP_API_URL}
ENV VUE_APP_SERVICE_NAME=${VUE_APP_SERVICE_NAME}
ENV VUE_APP_SERVICE_VERSION=${VUE_APP_SERVICE_VERSION}
ENV VUE_APP_OTEL_URL=${VUE_APP_OTEL_URL}

RUN yarn run build

FROM nginx:alpine as production-build

COPY ./.nginx/nginx.conf /etc/nginx/nginx.conf

## Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*

# Copy from the stag builder
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]