package repository

import (
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/connection"
"github.com/zhenisduissekov/http-3rdparty-task-service/internal/logger"
)

type Repository struct {
	conn *connection.Connection
	log  *logger.Logger
}

type DBConf struct {
	Host     string
	Port     string
	DB       string
	Username string
	Password string
}

func New(conn *connection.Connection, log *logger.Logger) *Repository {
	return &Repository{
		conn: conn,
		log:  log,
	}
}
