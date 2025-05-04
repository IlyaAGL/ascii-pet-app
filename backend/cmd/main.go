package main

import (
	"ascii_store/backend/internal/application"
	"ascii_store/backend/internal/infrastructure"
	"ascii_store/backend/internal/presentation/controllers"
)

func main() {
	repo := infrastructure.NewAsciiRepository()
	service := application.NewAsciiService(repo)
	handler := controllers.NewAsciiHandler(service)

	handler.StartApi()
}
