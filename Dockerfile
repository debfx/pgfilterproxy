FROM golang:1.22-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY pgbroker/go.mod ./pgbroker/
RUN go mod download

COPY *.go ./
COPY pgbroker/ ./pgbroker/
RUN go build -o /pgfilterproxy


FROM debian:bookworm-slim

COPY --from=builder /pgfilterproxy /pgfilterproxy

CMD ["/pgfilterproxy", "/config/pgfilterproxy.yaml"]
