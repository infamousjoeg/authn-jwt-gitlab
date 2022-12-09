FROM golang:1.19 as builder

WORKDIR /go/src/github.com/infamousjoeg/authn-jwt-gitlab
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM golang:1.19-alpine as builder-alpine

WORKDIR /go/src/github.com/infamousjoeg/authn-jwt-gitlab
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM ubuntu:kinetic as ubuntu

COPY --from=builder /go/bin/authn-jwt-gitlab /authn-jwt-gitlab

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

FROM alpine:3.17 as alpine

COPY --from=builder-alpine /go/bin/authn-jwt-gitlab /authn-jwt-gitlab

RUN apk add --no-cache ca-certificates && \
    update-ca-certificates

CMD ["/authn-jwt-gitlab"]

FROM redhat/ubi8:8.7 as ubi

COPY --from=builder /go/bin/authn-jwt-gitlab /authn-jwt-gitlab

CMD ["/authn-jwt-gitlab"]