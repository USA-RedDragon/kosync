FROM scratch

COPY --from=alpine:latest@sha256:71e9b7bae160ba167fa17414401b94cbcdaf3c32feeab9298bd794a2a0614c35 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ARG TARGETPLATFORM
COPY $TARGETPLATFORM/kosync /
ENTRYPOINT ["/kosync"]