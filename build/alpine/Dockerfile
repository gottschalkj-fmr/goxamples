#
# Copyright (c) 2020 gottschalkj-fmr.
# Licensed under the Apache License, Version 2.0
# http://www.apache.org/licenses/LICENSE-2.0
#
FROM golang:1.14-alpine AS builder
ARG MODULE
ENV MODULE=${MODULE}
ENV GOROOT=/usr/local/go
ENV CGO_ENABLED=0
ENV PATH=${GOROOT}/bin:${PATH}

WORKDIR /go/src
COPY go.mod go.sum ${MODULE}/
COPY cmd/ ${MODULE}/cmd/
COPY pkg/ ${MODULE}/pkg/
RUN go env && \
    cd ${MODULE} && \
    go test ./... && \
    go install ./...

FROM alpine:3.11 AS runtime
RUN addgroup -S appuser && \
    adduser -S -H -G appuser appuser && \
    mkdir -m 755 -p /app/bin && \
    chown -R appuser:appuser /app
COPY --from=builder --chown=appuser:appuser /go/bin/hello /app/bin/hello
EXPOSE 9090
WORKDIR /app
USER appuser
CMD ["bin/hello"]
