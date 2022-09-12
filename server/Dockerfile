FROM docker.bc-artifactory01/k8s-docker/golang:alpine3.16 AS builder
# FROM golang:alpine3.16 AS builder
WORKDIR /app
COPY *.go /app
COPY go.mod /app
RUN go build -o compote_server .

FROM docker.bc-artifactory01/k8s-docker/alpine:3.16.0
# FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder /app/compote_server /app/
ENTRYPOINT ["./compote_server"]

