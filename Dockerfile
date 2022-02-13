FROM golang:1.16-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

COPY src/*.go ./

COPY . .

RUN go build -o ./app ./src/api/main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .

EXPOSE 8000

ENTRYPOINT ["./app"]