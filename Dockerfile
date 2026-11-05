FROM scratch

COPY --from=alpine:latest@sha256:3f4d6b32e26703c6082839234f43769d5046b3ac7c910033911d1e71d8cd630a /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ARG TARGETPLATFORM
COPY $TARGETPLATFORM/kosync /
ENTRYPOINT ["/kosync"]