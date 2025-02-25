// @title SuperApp API
// @version 1.0
// @description API documentation for SuperApp
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"github.com/spashev/superapp/cmd/app"
	initializers "github.com/spashev/superapp/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	app := new(app.App)
	app.Run()
}
