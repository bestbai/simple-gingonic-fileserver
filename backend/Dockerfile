ARG GO_VERSION=1.11
FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
RUN mkdir -p /api
WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./simple_fileserver ./main.go

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /api /api/public
WORKDIR /api
COPY --from=builder /api/simple_fileserver .

EXPOSE 8081
ENTRYPOINT ["./simple_fileserver"]