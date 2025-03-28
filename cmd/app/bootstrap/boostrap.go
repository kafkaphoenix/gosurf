package bootstrap

import (
	"log/slog"

	"github.com/kafkaphoenix/gosurf/internal/repository"
)

// Run starts the application.
func Run() error {
	// 1. logger
	// 2. usecases
	// 3. handler
	// 4. server + routes
	// 5. gracefully shutdown
	slog.Info("A lot to do!") //nolint:sloglint //temporary

	fakedb, err := repository.NewFakeDB("db/actions.json", "db/users.json")
	if err != nil {
		return &AppError{Message: "error loading fakedb", Err: err}
	}

	totalUsers, totalActions := fakedb.GetTotal()
	slog.Info("loaded db", slog.Int("users count", totalUsers), slog.Int("actions count", totalActions)) //nolint:sloglint //temporary

	return nil
}
