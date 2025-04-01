// Package slogx provides extensions to the log/slog standard library.
package slogx

import (
	"log/slog"
	"os"
)

// New creates a new *[slog.Logger]. If debug is true, the returned logger will
// log at [slog.LevelDebug], otherwise it logs at [slog.LevelInfo]. Source
// lines are disabled to prevent displaying incorrect callers when using
// the logger in a [context.Context].
func New(debug bool) *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
	}

	var handler slog.Handler = slog.NewJSONHandler(os.Stdout, opts)
	if debug {
		opts.Level = slog.LevelDebug
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
