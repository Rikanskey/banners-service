FROM golang:alpine AS builder

WORKDIR /banner-service

COPY go.* ./

RUN go mod download && go mod verify

COPY ./. ./

RUN go build -v ./cmd/app

FROM alpine:3

COPY --from=builder /banner-service/app /banner-service/app
COPY --from=builder /banner-service/config/ /config

EXPOSE 8080

ENTRYPOINT ["/banner-service/app"]

