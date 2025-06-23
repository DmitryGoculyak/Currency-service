package main

import (
	"Currency-service/internal/db"
	"Currency-service/internal/service"
)

func main() {
	db.InitDB()
	service.RunService()
}
