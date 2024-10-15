FROM golang:1.23.2-alpine3.20 AS builder

COPY . /github.com/titoffon/chat-server/
WORKDIR /github.com/titoffon/chat-server/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/titoffon/chat-server/bin/chat_server .

CMD ["./chat_server"]