#FROM alpine:latest
FROM golang:alpine AS builder

WORKDIR /grain-acceptance

COPY go.* ./
#COPY build/package/dev ./
RUN ls

RUN go mod download && go mod verify

COPY ./. ./

#RUN #ls

RUN go build -v ./cmd/app

#RUN #ls ..

FROM alpine:3

#WORKDIR /grain-acceptance
COPY --from=builder /grain-acceptance/app /grain-acceptance/app
COPY --from=builder /grain-acceptance/config/ /config

#RUN #ls /grain-acceptance/
#RUN #ls /grain-acceptance/config
#ENV PORT 3000
EXPOSE 8080
#ENV GIN_MODE release

ENTRYPOINT ["/grain-acceptance/app"]