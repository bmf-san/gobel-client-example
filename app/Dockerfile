FROM golang:1.17.1-alpine

WORKDIR /go/gobel-client-example/app

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash && \
    go get -u github.com/cosmtrek/air && \
    go get -u golang.org/x/lint/golint

CMD air -c .air.toml
