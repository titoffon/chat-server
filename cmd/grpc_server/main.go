package main

import (
	"fmt"
	"log"
	"net"

	//"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/titoffon/chat-server/internal/config"
	"github.com/titoffon/chat-server/internal/server"
	"github.com/titoffon/chat-server/internal/storage"
	desc "github.com/titoffon/chat-server/pkg/chat_v1"
)

const grpcPort = 50051

/*type server struct {
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
func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Printf("Received message from %s: %s at %v\n", req.From, req.Text, req.Timestamp)
	log.Printf("Received message from %s: %s at %v\n", req.From, req.Text, req.Timestamp)
	return &emptypb.Empty{}, nil
}*/

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := storage.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	//Создаём экземпляр сервера
	srv := server.NewChatServiceServer(db)



	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()                        // создаём объект нового сервера
	reflection.Register(s)                       // включаем возможность сервера выдавать информацию о себе
	desc.RegisterChatServiceServer(s, srv)

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil { // запускаем сервер
		log.Fatalf("failed to serve: %v", err)
	}
}
