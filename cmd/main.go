package main

import (
	"superapp/cmd/app"
	initializers "superapp/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	app := new(app.App)
	app.Run()
}
