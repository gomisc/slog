package slog

import (
	"context"
)

type loggerKey struct{}

var dummy = newDummy() // nolint

// ToContext помещает логгер в контекст
func ToContext(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, log)
}

// MustFromContext возвращает логгер из контекста, в случае отсутствия
// логгера в контексте - паника
func MustFromContext(ctx context.Context) Logger {
	logger, ok := ctx.Value(loggerKey{}).(Logger)
	if !ok {
		return dummy
	}

	return logger
}
