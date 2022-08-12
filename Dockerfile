FROM golang:1.18.2 as builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/withings-prometheus-expoter \
    -ldflags '-s -w'

FROM alpine:latest as runner
RUN apk add curl
COPY settings.yaml access_token.json /
RUN chmod 766 access_token.json
COPY --from=builder /go/bin/withings-prometheus-expoter /app/withings-prometheus-expoter

# don't create homeDir.
RUN adduser -D -S -H exporter

USER exporter
ENTRYPOINT [ "/app/withings-prometheus-expoter" ]
