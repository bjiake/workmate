package logger

import (
	"errors"
	"log/slog"
	"os"
	"workmate/pkg/logger/slogpretty"
)

var unknownEnv = errors.New("unknown environment (should be local or prod)")

func New() *slog.Logger {
	var logger *slog.Logger

	logger = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	return logger
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
