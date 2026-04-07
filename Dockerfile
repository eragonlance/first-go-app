FROM golang:alpine AS builder

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd/appserver .

RUN go build -ldflags="-s -w" -trimpath -o app .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /workspace/app .

EXPOSE 8000

CMD ["./app"]
