FROM golang:1.17.1-alpine as builder

WORKDIR /go/gobel-client-example/app

ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

COPY . .

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o app

FROM --platform=$TARGETPLATFORM alpine

COPY --from=builder /go/gobel-client-example/app ./

CMD ["./app"]
