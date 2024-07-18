FROM golang:1.22-alpine3.20 AS builder

COPY . github.com/DenisCom3/m-auth
WORKDIR github.com/DenisCom3/m-auth

RUN go mod download
RUN go build -o ./bin/m-auth cmd/main.go

FROM alpine:3.20

WORKDIR /root/
COPY --from=builder /go/github.com/DenisCom3/m-auth/bin .
COPY --from=builder /go/github.com/DenisCom3/m-auth/.env .

CMD ["./m-auth", "--config-path", ".env" ]