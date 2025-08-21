package main

import (
	"fmt"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	repository := repository.NewRepository()
	service := service.NewService(repository)

	h := handler.NewHandler(service)
	h.RegisterRoutes(router)

	fmt.Println("Server starting on port 8080...")
	_ = router.Run(":8080")
}
