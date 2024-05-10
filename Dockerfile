FROM golang:1.22-alpine AS build-env
MAINTAINER Timofey Bakunin "bakunin.t@protonmail.com"

WORKDIR /app

RUN apk add --no-cache git make
COPY . .
RUN go mod download
RUN make build

FROM alpine as app
WORKDIR /app
COPY --from=build-env /app/exporter-merger /app/exporter-merger
ENTRYPOINT /app/exporter-merger
