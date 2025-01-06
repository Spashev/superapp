package main

import (
	"log"

	"github.com/joho/godotenv"

	"superapp/cmd/app"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

}

func main() {
	app := new(app.App)
	app.Run()
}
