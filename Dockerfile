FROM gcr.io/distroless/base

COPY bin/pass-output /pass-output

CMD ["/pass-output"]