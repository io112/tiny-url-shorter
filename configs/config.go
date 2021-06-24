package configs

import (
	"os"
)

var (
	PORT    = "8080"
	HOST    = "http://localhost"
	DB_NAME = "urldb.db"
)

func init() {
	PORT = getEnv("PORT", PORT)
	HOST = getEnv("HOST", HOST)
	DB_NAME = getEnv("DB_NAME", DB_NAME)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
