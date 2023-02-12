package store

import "github.com/jackc/pgx/v5/pgxpool"

type Store struct {
	db *pgxpool.Pool
}

func New(dsn string) *Store {
	return &Store{}
}
