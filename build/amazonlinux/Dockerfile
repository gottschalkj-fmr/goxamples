#
# Copyright (c) 2020 gottschalkj-fmr.
# Licensed under the Apache License, Version 2.0
# http://www.apache.org/licenses/LICENSE-2.0
#
FROM amazonlinux:2 AS builder
ARG MODULE
ENV MODULE=${MODULE}
ENV GOROOT=/usr/lib/golang
ENV GOPATH=/go
ENV GOCACHE=${GOPATH}/.cache/go-build
ENV GOENV=${GOPATH}/.config/go/env
ENV GO111MODULE=on
ENV PATH=${GOROOT}/bin:${PATH}

WORKDIR /go/src
COPY go.mod go.sum ${MODULE}/
COPY cmd/ ${MODULE}/cmd/
COPY pkg/ ${MODULE}/pkg/
RUN yum install -y golang && \
    go env && \
    cd ${MODULE} && \
    go test ./... && \
    go install ./...

FROM amazonlinux:2 AS runtime
RUN yum install -y shadow-utils && \
    useradd --system --no-create-home appuser && \
    mkdir -m 755 -p /app/bin && \
    chown -R appuser:appuser /app
COPY --from=builder --chown=appuser:appuser /go/bin/hello /app/bin/hello
EXPOSE 9090
WORKDIR /app
USER appuser
CMD ["bin/hello"]
