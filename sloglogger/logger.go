package sloglogger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	return slog.New(jsonHandler)
}
