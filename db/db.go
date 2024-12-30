package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	return &Database{Conn: db}, nil
}

func (db *Database) Close() {
	if db.Conn != nil {
		if err := db.Conn.Close(); err != nil {
			log.Printf("Error closing the database connection: %v", err)
		}
	}
}
