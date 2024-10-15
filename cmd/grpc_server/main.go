package main

import (
	"context"
	"fmt"
	"log"
	"net"

	//"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	//"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/titoffon/chat-server/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatServiceServer
}

// Create ...
func (s *server) Create(_ context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	fmt.Printf("Received Create chat request with users: %v\n", req.Usernames)
	log.Printf("Received Create chat request with users: %v", req.Usernames)
	return &desc.CreateChatResponse{Id: 1}, nil
}

// Delete ...
func (s *server) Delete(_ context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	fmt.Printf("Received Delete chat request for chat ID: %d\n", req.Id)
	log.Printf("Received Delete chat request for chat ID: %d\n", req.Id)
    return &emptypb.Empty{}, nil
}

// SendMessage ...
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
    fmt.Printf("Received message from %s: %s at %v\n", req.From, req.Text, req.Timestamp)
	log.Printf("Received message from %s: %s at %v\n", req.From, req.Text, req.Timestamp)
    return &emptypb.Empty{}, nil
}



func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()                        // создаём объект нового сервера
	reflection.Register(s)                       // включаем возможность сервера выдавать информацию о себе
	desc.RegisterChatServiceServer(s, &server{}) //второй параметр это структура, которая имплементировала API

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil { // запускаем сервер
		log.Fatalf("failed to serve: %v", err)
	}
}