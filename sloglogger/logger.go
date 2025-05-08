package sloglogger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	return slog.New(jsonHandler)

	//logger.Info("This is an Info message", slog.Int("version", 1.0))
	//logger.Debug("This is a Debug message")
	//logger.Info("This is an Info message")
	//logger.Warn("This is a Warning message")
	//logger.Error("This is an Error message")
}
