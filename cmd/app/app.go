package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"superapp/cmd/app/router"
	"superapp/config"
	"superapp/internal/db"
	"superapp/internal/util/token"
)

type App struct {
	fiberApp *fiber.App
}

func (a *App) Run() {
	cfg := config.NewConfig()
	jwtMaker := token.NewJWTMaker(cfg.JWTSecretKey)

	database, err := db.NewDatabase(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.Close()

	fmt.Printf("Server running on %s...\n", ":8080")

	a.fiberApp = router.RegisterRoutes(database.Conn, jwtMaker)

	if err := a.fiberApp.Listen(":8080"); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}

func (a *App) Shutdown() error {
	return a.fiberApp.Shutdown()
}
