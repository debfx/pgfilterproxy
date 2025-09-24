FROM golang:1.25.1-trixie AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY pgbroker/go.mod ./pgbroker/
RUN go mod download

COPY *.go ./
COPY pgbroker/ ./pgbroker/
RUN go build -o /pgfilterproxy


FROM debian:trixie-slim

COPY --from=builder /pgfilterproxy /pgfilterproxy

CMD ["/pgfilterproxy", "/config/pgfilterproxy.yaml"]
