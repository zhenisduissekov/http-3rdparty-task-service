package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type AuthConf struct {
	Username string
	Password string
}

type DBConf struct {
	Host     string
	Port     string
	Db       string
	Username string
	Password string
}

type Connection struct {
	Host     string
	Port     string
	Db       string
	Username string
	Password string
}

type GoogleConf struct {
	Url string
}

type BingConf struct {
	Url string
}

type ChatGPTConf struct {
	Url   string
	Token string
}

type Conf struct {
	ServiceName string
	Auth        AuthConf
	Google      GoogleConf
	ChatGPT     ChatGPTConf
	Bing        BingConf
	Connection  Connection
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
		Google: GoogleConf{
			Url: os.Getenv("GOOGLE_API_URL"),
		},
		ChatGPT: ChatGPTConf{
			Url: os.Getenv("CHAT_GPT_URL"),
		},
		Bing: BingConf{
			Url: os.Getenv("BING_API_URL"),
		},
		Connection: Connection{
			Host:     os.Getenv("POSTGRE_HOST"),
			Port:     os.Getenv("POSTGRE_PORT"),
			Db:       os.Getenv("POSTGRE_DB"),
			Username: os.Getenv("POSTGRE_USERNAME"),
			Password: os.Getenv("POSTGRE_PASSWORD"),
		},
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}
