package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(databaseURL string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), databaseURL)
	
	if err != nil {
		return nil, fmt.Errorf("erro ao criar pool de conexão: %w", err)
	}

	// Garante que a conexão realmente funciona.
	if err := db.Ping(context.Background()); err != nil {
		db.Close()
		return nil, fmt.Errorf("erro ao conectar no banco: %w", err)
	}

	return db, nil
}