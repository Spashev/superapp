package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"superapp/cmd/app/router"
	"superapp/config"
	"superapp/internal/db"
)

type App struct {
	http *http.Server
}

func (a *App) Run() {
	cfg := config.NewConfig()

	database, err := db.NewDatabase(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.Close()

	fmt.Printf("Server running on %s...\n", ":8080")

	r := router.RegisterRoutes(database.Conn)
	a.http = &http.Server{
		Addr:           ":8080",
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        r,
	}
	if err := a.http.ListenAndServe(); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}

func (a *App) Shotdown(ctx context.Context) error {
	return a.http.Shutdown(ctx)
}
