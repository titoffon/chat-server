package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/titoffon/chat-server/internal/config"
)

func NewDB(cfg *config.Config) (*pgxpool.Pool, error) {
    dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
        cfg.Database.User,
        cfg.Database.Password,
        cfg.Database.Host,
        cfg.Database.Port,
        cfg.Database.DBName,
    )

    pool, err := pgxpool.Connect(context.Background(), dsn)
    if err != nil {
        return nil, fmt.Errorf("unable to connect to database: %w", err)
    }

    return pool, nil
}
