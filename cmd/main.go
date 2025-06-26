package main

import (
	"Currency-service/internal/container"
)

func main() {
	app := container.Build()

	app.Run()
}
