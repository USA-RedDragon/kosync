FROM scratch

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ARG TARGETPLATFORM
COPY $TARGETPLATFORM/kosync /
ENTRYPOINT ["/kosync"]