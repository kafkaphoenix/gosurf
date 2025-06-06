package bootstrap

import (
	"log/slog"
	"os"

	"github.com/kafkaphoenix/gosurf/internal/repository/db"
	"github.com/kafkaphoenix/gosurf/internal/repository/server"
	"github.com/kafkaphoenix/gosurf/internal/usecases"
)

// Run starts the application.
func Run() error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := db.NewFakeDB("db/users.json", "db/actions.json")
	if err != nil {
		return &AppError{Message: "error loading fakedb", Err: err}
	}

	userService := usecases.NewUserService(db)
	userHandler := server.NewUserHandler(userService)
	actionService := usecases.NewActionService(db)
	actionHandler := server.NewActionHandler(actionService)

	srv, err := server.New(logger)
	if err != nil {
		return &AppError{
			Message: "error creating HTTP server",
			Err:     err,
		}
	}

	if err := srv.RegisterRoutes(userHandler.RegisterRoutes, actionHandler.RegisterRoutes); err != nil {
		return &AppError{
			Message: "error registering server routes",
			Err:     err,
		}
	}

	if err := srv.Start(); err != nil {
		return &AppError{
			Message: "error starting HTTP server",
			Err:     err,
		}
	}

	return nil
}
