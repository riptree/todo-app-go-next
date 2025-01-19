package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"task-management/internal/infrastructure/db"
	"task-management/internal/interface/router"
	"task-management/internal/package/config"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func run() error {
	err := config.Init()
	if err != nil {
		return errors.Wrap(err, "failed to initialize config")
	}

	db, err := db.OpenDB()
	if err != nil {
		return errors.Wrap(err, "failed to initialize a new database")
	}

	e := echo.New()

	router.Init(e, db)

	port := "8080"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: e,
	}

	slog.Info(fmt.Sprintf("listening on port %s", port))
	if err := srv.ListenAndServe(); err != nil {
		return errors.Wrap(err, "failed to listen and serve")
	}

	return nil
}
