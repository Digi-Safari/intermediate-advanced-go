package main

import (
	"log/slog"
	"os"
)

func setupSlog() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}
