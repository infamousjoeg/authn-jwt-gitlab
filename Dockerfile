FROM golang:1.19 as builder

WORKDIR /go/src/github.com/infamousjoeg/authn-jwt-gitlab
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM ubuntu:jammy

COPY --from=builder /go/bin/authn-jwt-gitlab /authn-jwt-gitlab

CMD ["/authn-jwt-gitlab"]