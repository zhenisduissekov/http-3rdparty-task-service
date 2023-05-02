package logger

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rzajac/zltest"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
)

const (
	msg = "message"
)

// testing different levels of logging: error, debug, info, warn

func TestLogger_Error(t *testing.T) {
	tst := zltest.New(t)
	log := zerolog.New(tst).With().Timestamp().Logger()
	srv := New(&config.Conf{LogLevel: "info"})
	srv.logger = &log
	srv.Error().Msg(msg)
	ent := tst.LastEntry()
	ent.ExpMsg(msg)
	ent.ExpLevel(zerolog.ErrorLevel)
}

func TestLogger_Debug(t *testing.T) {
	tst := zltest.New(t)
	log := zerolog.New(tst).With().Timestamp().Logger()
	srv := New(&config.Conf{LogLevel: "info"})
	srv.logger = &log
	srv.Debug().Msg(msg)
	ent := tst.LastEntry()
	ent.ExpMsg(msg)
	ent.ExpLevel(zerolog.DebugLevel)
}

func TestLogger_Info(t *testing.T) {
	tst := zltest.New(t)
	log := zerolog.New(tst).With().Timestamp().Logger()
	srv := New(&config.Conf{LogLevel: "info"})
	srv.logger = &log
	srv.Info().Msg(msg)
	ent := tst.LastEntry()
	ent.ExpMsg(msg)
	ent.ExpLevel(zerolog.InfoLevel)
}

func TestLogger_Warn(t *testing.T) {
	tst := zltest.New(t)
	log := zerolog.New(tst).With().Timestamp().Logger()
	srv := New(&config.Conf{LogLevel: "info"})
	srv.logger = &log
	srv.Warn().Msg(msg)
	ent := tst.LastEntry()
	ent.ExpMsg(msg)
	ent.ExpLevel(zerolog.WarnLevel)
}

func TestLogger_ErrorExists(t *testing.T) {
	tst := zltest.New(t)
	log := zerolog.New(tst).With().Timestamp().Logger()
	srv := New(&config.Conf{LogLevel: "info"})
	srv.logger = &log
	srv.Error().Err(errors.New("test")).Msg(msg)
	ent := tst.LastEntry()
	ent.ExpMsg(msg)
	ent.ExpLevel(zerolog.ErrorLevel)
	ent.ExpErr(errors.New("test"))
}
