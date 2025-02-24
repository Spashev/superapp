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
