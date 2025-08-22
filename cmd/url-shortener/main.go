package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"url-shortener/internal/config"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	router := gin.Default()

	repo, err := repository.NewRedisRepository(cfg.RedisAddr)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	service := service.NewService(repo)
	h := handler.NewHandler(service)
	h.RegisterRoutes(router)

	log.Printf("Server starting on port %s...", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
