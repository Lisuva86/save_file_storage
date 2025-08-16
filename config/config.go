package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)
const (
	StorageRoot = "./storage"
	EnvDev      = "dev"
	EnvProd     = "prod"
	URL         = "https://test_med2/"
)
const (
	TokenExpiry  = 24 * time.Hour
)
func GetAdminUsername() string {
	user := os.Getenv("ADMIN_USERNAME")
	if user == "" {
		return "admin" // fallback
	}
	return user
}

func GetAdminPassword() string {
	pass := os.Getenv("ADMIN_PASSWORD")
	if pass == "" {
		log.Fatal("ADMIN_PASSWORD is not set")
	}
	return pass
}
func init() {
	// Загружаем .env только в dev
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Println("Не удалось загрузить .env файл")
		}
	}
}

// GetJWTSecret возвращает секретный ключ из переменной окружения
func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("ОШИБКА: JWT_SECRET_KEY не установлен в переменных окружения!")
	}
	return secret
}

// GetPort возвращает порт (по умолчанию 8080)
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}