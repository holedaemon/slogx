package slogx

import (
	"context"
	"io"
	"log/slog"
)

type loggerKey struct{}

var nopLogger = slog.New(slog.NewTextHandler(io.Discard, nil))

// FromContext returns a *[slog.Logger] from the given [context.Context]. If the
// context does not hold a logger, then a nop logger is returned.
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		return logger
	}

	return nopLogger
}

// WithLogger returns a [context.Context] with the given logger added to its
// values.
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// Debug logs at the debug level to the logger held by the given
// [context.Context]
func Debug(ctx context.Context, msg string, attr ...any) {
	FromContext(ctx).DebugContext(ctx, msg, attr...)
}

// Info logs at the info level to the logger held by the given
// [context.Context].
func Info(ctx context.Context, msg string, attr ...any) {
	FromContext(ctx).InfoContext(ctx, msg, attr...)
}

// Warn logs at the warn level to the logger held by the given
// [context.Context].
func Warn(ctx context.Context, msg string, attr ...any) {
	FromContext(ctx).WarnContext(ctx, msg, attr...)
}

// Error logs at the error level to the logger held by the given
// [context.Context].
func Error(ctx context.Context, msg string, attr ...any) {
	FromContext(ctx).ErrorContext(ctx, msg, attr...)
}

// With appends the given attributes to the logger held by the given
// [context.Context].
func With(ctx context.Context, attr ...any) context.Context {
	return WithLogger(ctx, FromContext(ctx).With(attr...))
}

// WithGroup adds the given group to the logger held by the given
// [context.Context].
func WithGroup(ctx context.Context, name string) context.Context {
	return WithLogger(ctx, FromContext(ctx).WithGroup(name))
}
