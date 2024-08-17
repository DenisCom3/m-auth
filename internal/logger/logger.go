package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

const (
	envDev  = "dev"
	envProd = "prod"

	logPath = "./storage/logs/app.log"
)

var globalLogger *slog.Logger

func Info(msg string, args ...any) {
	if globalLogger == nil {
		return
	}
	fArgs := formatArgs(args...)
	globalLogger.Info(msg, fArgs...)
}

func Error(msg string, args ...any) {
	if globalLogger == nil {
		return
	}
	fArgs := formatArgs(args...)
	globalLogger.Error(msg, fArgs...)
}

func Debug(msg string, args ...any) {
	if globalLogger == nil {
		return
	}
	fArgs := formatArgs(args...)
	globalLogger.Debug(msg, fArgs...)
}

func Init(env string) error {
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Printf("failed to open log file: %v", err)
		return err
	}

	w := io.Writer(logFile)

	switch env {
	case envDev:
		globalLogger = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		globalLogger = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		globalLogger = slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return nil
}

func errMarshal(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
