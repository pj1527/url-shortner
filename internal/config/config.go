package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port      string
	RedisAddr string
}

func LoadConfig() AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	return AppConfig{
		Port:      port,
		RedisAddr: redisAddr,
	}
}
