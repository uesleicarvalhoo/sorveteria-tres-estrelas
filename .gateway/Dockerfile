FROM kong:3.1.1-alpine

USER root

RUN apk add --no-cache git \
    && luarocks install kong-jwt2header
