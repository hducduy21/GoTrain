package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configs struct {
	PublicHost string
	Port       string
	DbConnStr  string
}

var Envs = initConfig()

func initConfig() *Configs {
	godotenv.Load()

	return &Configs{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DbConnStr:  getEnv("DB_CONN_STR", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
