package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
)

var (
	log   *slog.Logger
	debug bool
)

// simpleFormatter is a custom handler for clean text output
type simpleFormatter struct{}

func (f *simpleFormatter) Handle(_ context.Context, r slog.Record) error {
	// Just print the message with attributes in readable format
	fmt.Print(r.Message)
	if r.NumAttrs() > 0 {
		fmt.Print(" ")
		r.Attrs(func(a slog.Attr) bool {
			fmt.Printf("%v=%v ", a.Key, a.Value)
			return true
		})
	}
	fmt.Println()
	return nil
}

func (f *simpleFormatter) Enabled(_ context.Context, level slog.Level) bool {
	return level >= slog.LevelInfo
}

func (f *simpleFormatter) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &simpleFormatter{}
}

func (f *simpleFormatter) WithGroup(name string) slog.Handler {
	return &simpleFormatter{}
}

// Init initializes the logger with the given debug flag
func Init(isDebug bool) {
	debug = isDebug

	if debug {
		// Debug mode: structured slog output
		opts := &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}
		handler := slog.NewTextHandler(os.Stdout, opts)
		log = slog.New(handler)
	} else {
		// Normal mode: simple human-readable text output
		handler := &simpleFormatter{}
		log = slog.New(handler)
	}
}

// Debug logs at debug level with file location
func Debug(msg string, args ...any) {
	if !debug {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if ok {
		idx := strings.LastIndex(file, "\\")
		if idx == -1 {
			idx = strings.LastIndex(file, "/")
		}
		if idx != -1 {
			file = file[idx+1:]
		}
		args = append([]any{"location", fmt.Sprintf("%s:%d", file, line)}, args...)
	}
	log.Debug(msg, args...)
}

// Info logs at info level
func Info(msg string, args ...any) {
	log.Info(msg, args...)
}

// Verbose logs only in debug mode
func Verbose(msg string, args ...any) {
	if !debug {
		return
	}
	log.Info(msg, args...)
}

// Error logs at error level
func Error(msg string, err error, args ...any) {
	args = append([]any{"error", err}, args...)
	log.Error(msg, args...)
}

// Warn logs at warn level
func Warn(msg string, args ...any) {
	log.Warn(msg, args...)
}

// With returns a child logger with additional context
func With(args ...any) *slog.Logger {
	return log.With(args...)
}

// GetLogger returns the underlying slog.Logger
func GetLogger() *slog.Logger {
	return log
}

// IsDebug returns whether debug mode is enabled
func IsDebug() bool {
	return debug
}
