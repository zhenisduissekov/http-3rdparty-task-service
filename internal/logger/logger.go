package logger

import (
	"os"
		"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
)

type Logger struct {
	logger *zerolog.Logger
}

type Log interface {
	Debug() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
	Trace() *zerolog.Event
	Panic() *zerolog.Event
}

func New(cnf *config.Conf) *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Logger = log.With().Caller().Logger()
	logLevel, err := zerolog.ParseLevel(strings.ToLower(cnf.LogLevel))

	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	log.WithLevel(logLevel).Msgf("Log level set to %s", logLevel)
	zerolog.SetGlobalLevel(logLevel)

	return &Logger{
		logger: &log.Logger,
	}
}

func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}
func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}
func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}
func (l *Logger) Trace() *zerolog.Event {
	return l.logger.Trace()
}

func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic()
}