FROM kong:3

USER root

RUN apk add --no-cache git \
    && luarocks install kong-jwt2header

USER kong
