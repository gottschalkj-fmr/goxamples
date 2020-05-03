#
# Copyright (c) 2020 krautbax.
# Licensed under the Apache License, Version 2.0
# http://www.apache.org/licenses/LICENSE-2.0
#
FROM golang:1.14-buster
ARG MODULE
ENV MODULE=${MODULE}
ENV GOROOT=/usr/local/go
ENV PATH="${GOROOT}/bin:${PATH}"

WORKDIR /go/src
COPY go.mod go.sum ${MODULE}/
COPY cmd/ ${MODULE}/cmd/
COPY pkg/ ${MODULE}/pkg/
RUN go env && \
    cd ${MODULE} && \
    go test ./... && \
    go install ./...

EXPOSE 9090
CMD ["/go/bin/hello"]
