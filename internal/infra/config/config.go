package config

import (
	"os"
	"strconv"

	"github.com/DevAntonioJorge/go-notes/pkg/logger"
)

type AppConfig struct {
	Port        string
	DBUrl       string
	MongoDBUrl  string
	JWTSecret   string
	MetricsPort string
	LogLevel    logger.LogLevel
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func GetEnvInt(key string, fallback int) logger.LogLevel {
	value := os.Getenv(key)
	if value == "" {
		return logger.LogLevel(fallback)
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return logger.LogLevel(fallback)
	}
	return logger.LogLevel(intValue)
}

func GetConfig() *AppConfig {
	return &AppConfig{
		Port:        GetEnv("PORT", ":8000"),
		DBUrl:       GetEnv("DATABASE_URL", "postgresql://postgres:password@host:5432/postgres"),
		MongoDBUrl:  GetEnv("MONGO_URL", "mongodb://localhost:27017"),
		JWTSecret:   GetEnv("JWT_SECRET", "wegboipnioncI[ONCV9EWBNVG98A8WBGF3Q]"),
		MetricsPort: GetEnv("METRICS_PORT", ":8001"),
		LogLevel:    GetEnvInt("LOG_LEVEL", 1),
	}
}
