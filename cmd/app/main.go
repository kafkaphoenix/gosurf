package main

import (
	"log/slog"

	"github.com/kafkaphoenix/gosurf/cmd/app/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		slog.Error("Application failed to run", slog.String("error", err.Error()))
	}
}
