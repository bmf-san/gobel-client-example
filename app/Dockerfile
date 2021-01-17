FROM golang:1.13.0-alpine

ENV SERVER_PORT=8081
ENV HTTP_API_SCHEME=http
ENV HTTP_API_HOST=gobel-api:8080
ENV TIME_ZONE=Asia/Tokyo

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash && \
    go get github.com/oxequa/realize && \
    go get -u golang.org/x/lint/golint

WORKDIR /go/src/github.com/bmf-san/gobel-client-example/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

