FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

FROM scratch

COPY vv /
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/vv"]
