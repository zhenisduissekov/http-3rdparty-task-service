package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type AuthConf struct {
	Username string
	Password string
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
			Username: os.Getenv("USER_NAME_API"),
			Password: os.Getenv("PASSWORD_API"),
		},
		Cache: CacheConf{
			DefaultExpiration: getEnvDuration("CACHE_DEFAULT_EXPIRATION", "60s"),
			CleanupInterval:   getEnvDuration("CACHE_CLEANUP_INTERVAL", "120s"),
		},
		LogLevel: os.Getenv("LOG_LEVEL"),
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
