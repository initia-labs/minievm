package jsonrpc

import (
	"context"
	"log/slog"

	"cosmossdk.io/log"

	ethlog "github.com/ethereum/go-ethereum/log"
)

var _ slog.Handler = (*logHandler)(nil)

type logHandler struct {
	log.Logger
	group string
	attrs []slog.Attr
}

func newLogger(logger log.Logger) *logHandler {
	return &logHandler{Logger: logger}
}

// Enabled implements slog.Handler.
func (l *logHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

// Handle implements slog.Handler.
func (l *logHandler) Handle(ctx context.Context, r slog.Record) error {
	attrs := make([]any, 2*len(l.attrs))
	for i, attr := range l.attrs {
		attrs[i*2] = attr.Key
		attrs[i*2+1] = attr.Value
	}

	switch r.Level {
	case ethlog.LevelTrace, ethlog.LevelDebug:
		l.Logger.Debug(r.Message, attrs...)
	case ethlog.LevelInfo, ethlog.LevelWarn:
		l.Logger.Info(r.Message, attrs...)
	case ethlog.LevelError, ethlog.LevelCrit:
		l.Logger.Error(r.Message, attrs...)
	}

	return nil
}

// WithAttrs implements slog.Handler.
func (l *logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &logHandler{Logger: l.Logger, group: l.group, attrs: append(l.attrs, attrs...)}
}

// WithGroup implements slog.Handler.
func (l *logHandler) WithGroup(name string) slog.Handler {
	return &logHandler{Logger: l.Logger, group: name, attrs: l.attrs}
}
