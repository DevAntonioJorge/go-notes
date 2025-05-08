package main

import "os"

type AppConfig struct {
	Port string
	DBUrl string
	JWTSecret string
	MetricsPort string
}

func GetEnv(key, fallback string) string{
	value := os.Getenv(key)
	if value == ""{
		return fallback
	}
	return value
}

func GetConfig() *AppConfig{
	return &AppConfig{
		Port: GetEnv("PORT", ":8000"),
		DBUrl: GetEnv("DATABASE_URL", "postgresql://postgres:password@host:5432/postgres"),
		JWTSecret: GetEnv("JWT_SECRET", "wegboipnioncI[ONCV9EWBNVG98A8WBGF3Q]"),
		MetricsPort: GetEnv("METRICS_PORT", ":8001"),
	}
}