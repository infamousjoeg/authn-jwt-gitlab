FROM golang:1.19 as builder

WORKDIR /go/src/github.com/infamousjoeg/authn-jwt-gitlab
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM gitlab/gitlab-runner:latest as ubuntu

COPY --from=builder /go/bin/authn-jwt-gitlab /authn-jwt-gitlab

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/*

CMD ["/authn-jwt-gitlab"]

FROM gitlab/gitlab-runner:alpine as alpine

COPY --from=builder /go/bin/authn-jwt-gitlab /authn-jwt-gitlab

RUN apk add --no-cache ca-certificates && \
    update-ca-certificates

CMD ["/authn-jwt-gitlab"]

FROM gitlab/gitlab-runner:ubi-fips as ubi

COPY --from=builder /go/bin/authn-jwt-gitlab /authn-jwt-gitlab

RUN yum install -y ca-certificates && \
    update-ca-certificates && \
    yum clean all

CMD ["/authn-jwt-gitlab"]