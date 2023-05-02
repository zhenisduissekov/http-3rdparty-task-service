package logger

import (
"testing"

"github.com/pkg/errors"
"github.com/rs/zerolog"
"github.com/rzajac/zltest"
"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
)

func TestLogger_Error(t *testing.T) {
	tst := zltest.New(t)

	// Configure zerolog and pas tester as a writer.
	log := zerolog.New(tst).With().Timestamp().Logger()

	// Inject log to tested service or package.
	srv := New(&config.Conf{LogLevel: "info"})
	srv.logger = &log
	// --- When ---
	srv.Error().Msg("message")
	// Test if log messages were generated properly.
	ent := tst.LastEntry()
	ent.ExpMsg("message")
	ent.ExpLevel(zerolog.ErrorLevel)
}

func TestLogger_Debug(t *testing.T) {
	tst := zltest.New(t)

	// Configure zerolog and pas tester as a writer.
	log := zerolog.New(tst).With().Timestamp().Logger()

	// Inject log to tested service or package.
	srv := New(&config.Conf{LogLevel: "info"})
	srv.logger = &log
	// --- When ---
	srv.Debug().Msg("message")
	// Test if log messages were generated properly.
	ent := tst.LastEntry()
	ent.ExpMsg("message")
	ent.ExpLevel(zerolog.DebugLevel)
}

func TestLogger_Info(t *testing.T) {
	tst := zltest.New(t)

	// Configure zerolog and pas tester as a writer.
	log := zerolog.New(tst).With().Timestamp().Logger()

	// Inject log to tested service or package.
	srv := New(&config.Conf{LogLevel: "info"})
	srv.logger = &log
	// --- When ---
	srv.Info().Msg("message")
	// Test if log messages were generated properly.
	ent := tst.LastEntry()
	ent.ExpMsg("message")
	ent.ExpLevel(zerolog.InfoLevel)
}

func TestLogger_Warn(t *testing.T) {
	tst := zltest.New(t)

	// Configure zerolog and pas tester as a writer.
	log := zerolog.New(tst).With().Timestamp().Logger()

	// Inject log to tested service or package.
	srv := New(&config.Conf{LogLevel: "info"})
	srv.logger = &log
	// --- When ---
	srv.Warn().Msg("message")
	// Test if log messages were generated properly.
	ent := tst.LastEntry()
	ent.ExpMsg("message")
	ent.ExpLevel(zerolog.WarnLevel)
}

func TestLogger_ErrorExists(t *testing.T) {
	tst := zltest.New(t)

	// Configure zerolog and pas tester as a writer.
	log := zerolog.New(tst).With().Timestamp().Logger()

	// Inject log to tested service or package.
	srv := New(&config.Conf{LogLevel: "info"})
	srv.logger = &log
	// --- When ---
	srv.Error().Err(errors.New("test")).Msg("message")
	// Test if log messages were generated properly.
	ent := tst.LastEntry()
	ent.ExpMsg("message")
	ent.ExpLevel(zerolog.ErrorLevel)
	ent.ExpErr(errors.New("test"))
}
