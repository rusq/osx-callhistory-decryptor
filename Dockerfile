FROM golang:1.15
LABEL maintainer="github.com/rusq"

COPY . /src

WORKDIR /src

RUN go test ./...
RUN go build