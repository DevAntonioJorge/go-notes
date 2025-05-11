package main

import "os"

type AppConfig struct {
	Port        string
	DBUrl       string
	MongoDBUrl  string
	JWTSecret   string
	MetricsPort string
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func GetConfig() *AppConfig {
	return &AppConfig{
		Port:        GetEnv("PORT", ":8000"),
		DBUrl:       GetEnv("DATABASE_URL", "postgresql://postgres:password@host:5432/postgres"),
		MongoDBUrl:  GetEnv("MONGO_URL", "mongodb://localhost:27017"),
		JWTSecret:   GetEnv("JWT_SECRET", "wegboipnioncI[ONCV9EWBNVG98A8WBGF3Q]"),
		MetricsPort: GetEnv("METRICS_PORT", ":8001"),
	}
}
