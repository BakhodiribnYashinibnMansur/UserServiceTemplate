package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/spf13/cast"
)

// Config struct
type Config struct {
	PostgresHost           string
	PostgresPort           int
	PostgresDatabase       string
	PostgresUser           string
	PostgresPassword       string
	PostgresSSLMode        string
	PostgresMaxConnections int32
	ServerHost             string
	ServerPort             string
}

// load for loading a config
func load() *Config {
	return &Config{
		PostgresHost:           cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:           cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresDatabase:       cast.ToString(getOrReturnDefault("POSTGRES_DB", "")),
		PostgresUser:           cast.ToString(getOrReturnDefault("POSTGRES_USER", "")),
		PostgresPassword:       cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "")),
		PostgresSSLMode:        cast.ToString(getOrReturnDefault("POSTGRES_SSLMODE", "disable")),
		PostgresMaxConnections: cast.ToInt32(getOrReturnDefault("POSTGRES_MAX_CONNECTIONS", "30")),
		ServerHost:             cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost")),
		ServerPort:             cast.ToString(getOrReturnDefault("SERVER_PORT", ":9000")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

var (
	instance *Config
	once     sync.Once
)

// Get ...
func Get() *Config {
	once.Do(func() {
		instance = load()
	})

	return instance
}
