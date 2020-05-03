#
# Copyright (c) 2020 krautbax.
# Licensed under the Apache License, Version 2.0
# http://www.apache.org/licenses/LICENSE-2.0
#
FROM amazonlinux:2018.03.0.20200318.1
ARG MODULE
ENV MODULE=${MODULE}
ENV GOROOT=/usr/lib/golang
ENV GOPATH=/app/go
ENV GOCACHE=${GOPATH}/.cache/go-build
ENV GOENV=${GOPATH}/.config/go/env
ENV GO111MODULE=on
ENV PATH=${GOROOT}/bin:${PATH}

RUN yum install -y golang && \
    useradd --system --no-create-home appuser && \
    mkdir -m 755 -p /app/go/src && \
    chown -R appuser:appuser /app

WORKDIR /app/go/src
COPY --chown=appuser:appuser go.mod go.sum ${MODULE}/
COPY --chown=appuser:appuser cmd/ ${MODULE}/cmd/
COPY --chown=appuser:appuser pkg/ ${MODULE}/pkg/

USER appuser
RUN go env && \
    cd ${MODULE} && \
    go test ./... && \
    go install ./...

EXPOSE 9090
CMD ["/app/go/bin/hello"]
