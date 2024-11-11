package server

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	desc "github.com/titoffon/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

type ChatServiceServer struct {
    desc.UnimplementedChatServiceServer
    db *pgxpool.Pool
}

func NewChatServiceServer(db *pgxpool.Pool) *ChatServiceServer {
    return &ChatServiceServer{
        db: db,
    }
}

// Create ...
func (s *ChatServiceServer) Create(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
    log.Printf("Received Create chat request with users: %v", req.Usernames)

    // Сохраняем чат в базе данных
    var chatID int64
    query := `INSERT INTO chats (usernames) VALUES ($1) RETURNING id`
    err := s.db.QueryRow(ctx, query, req.Usernames).Scan(&chatID)
    if err != nil {
        return nil, fmt.Errorf("failed to create chat: %w", err)
    }

    return &desc.CreateChatResponse{Id: chatID}, nil
}

// Delete ...
func (s *ChatServiceServer) Delete(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
    log.Printf("Received Delete chat request for chat ID: %d", req.Id)

    // Удаляем чат из базы данных
    query := `DELETE FROM chats WHERE id = $1`
    _, err := s.db.Exec(ctx, query, req.Id)
    if err != nil {
        return nil, fmt.Errorf("failed to delete chat: %w", err)
    }

    return &emptypb.Empty{}, nil
}

// SendMessage ...
func (s *ChatServiceServer) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
    log.Printf("Received message from %s: %s at %v", req.From, req.Text, req.Timestamp)

    // Сохраняем сообщение в базе данных
    query := `INSERT INTO messages (chat_id, sender, text, timestamp) VALUES ($1, $2, $3, $4)`
    _, err := s.db.Exec(ctx, query, req.ChatId, req.From, req.Text, req.Timestamp.AsTime())
    if err != nil {
        return nil, fmt.Errorf("failed to send message: %w", err)
    }

    return &emptypb.Empty{}, nil
}
