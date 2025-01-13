package db

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Conn *sqlx.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(60 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	return &Database{Conn: db}, nil
}

func (db *Database) Close() {
	if err := db.Conn.Close(); err != nil {
		fmt.Printf("Error closing the database: %v\n", err)
	}
}
