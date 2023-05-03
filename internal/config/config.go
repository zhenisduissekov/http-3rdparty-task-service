package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type AuthConf struct {
	Username       string
	Password       string
	Port           string
	RequestTimeout time.Duration
}

type CacheConf struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

type Conf struct {
	ServiceName string
	Auth        AuthConf
	Cache       CacheConf
	LogLevel    string
}

func New() *Conf {
	err := godotenv.Load()
	if err != nil {
		log.Err(err).Msg("error loading env file")
	}
	return &Conf{
		ServiceName: "http_3rdparty_task_service",
		Auth: AuthConf{
			Username:       getEnv("USER_NAME_API", "admin"),
			Password:       getEnv("PASSWORD_API", "admin"),
			Port:           getEnv("PORT_API", "3000"),
			RequestTimeout: getEnvDuration("REQUEST_TIMEOUT", "10s"),
		},
		Cache: CacheConf{
			DefaultExpiration: getEnvDuration("CACHE_DEFAULT_EXPIRATION", "60s"),
			CleanupInterval:   getEnvDuration("CACHE_CLEANUP_INTERVAL", "120s"),
		},
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

func getEnvDuration(key string, defaultValue string) time.Duration {
	value, ok := os.LookupEnv(key)
	if !ok {
		dur, err := time.ParseDuration(defaultValue)
		if err != nil {
			log.Err(err).Msg("error parsing duration")
		}
		return dur
	}
	dur, err := time.ParseDuration(value)
	if err != nil {
		log.Err(err).Msg("error parsing duration")
	}
	return dur
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
