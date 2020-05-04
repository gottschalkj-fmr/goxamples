#
# Copyright (c) 2020 krautbax.
# Licensed under the Apache License, Version 2.0
# http://www.apache.org/licenses/LICENSE-2.0
#
FROM golang:1.14-buster AS builder
ARG MODULE
ENV MODULE=${MODULE}
ENV GOROOT=/usr/local/go
ENV PATH=${GOROOT}/bin:${PATH}

WORKDIR /go/src
COPY go.mod go.sum ${MODULE}/
COPY cmd/ ${MODULE}/cmd/
COPY pkg/ ${MODULE}/pkg/
RUN go env && \
    cd ${MODULE} && \
    go test ./... && \
    go install ./...

# Could use ubuntu:bionic, for example, since that is Debian-based as well.
FROM debian:10.3-slim AS runtime
RUN useradd --system --no-create-home appuser && \
    mkdir -m 755 -p /app/bin && \
    chown -R appuser:appuser /app
COPY --from=builder --chown=appuser:appuser /go/bin/hello /app/bin/hello
EXPOSE 9090
USER appuser
CMD ["/app/bin/hello"]
