FROM golang:alpine AS builder

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd/appserver .

RUN --mount=type=cache,target=~/.cache/go-build \
    go build -o app .

FROM scratch

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /workspace/app .

EXPOSE 8000

CMD ["./app"]