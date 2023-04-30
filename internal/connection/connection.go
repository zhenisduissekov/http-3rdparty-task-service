package connection

import (
	"context"
	"fmt"
"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/logger"
)

const (
	sslMode = "disable"
	timeOut = 15 * time.Second
)

type Connection struct {
	pool *pgxpool.Pool
	log  *logger.Logger
}

func New(cn *config.Conf, log *logger.Logger) *Connection {
	cnf := cn.Connection
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s application_name=%s sslmode=%s",
		cnf.Host, cnf.Port, cnf.Username, cnf.Password, cnf.Db, cn.ServiceName, sslMode) //good to know application name for debugging in pgadmin
	log.Trace().Msgf("connection string %s", connectionString)

	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatal().Err(err).Msg("error during connecting to DB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("error during pinging to DB")
	}

	return &Connection{
		pool: pool,
		log:  log,
	}
}
