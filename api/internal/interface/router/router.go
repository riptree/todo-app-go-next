package router

import (
	"todo-app/internal/infrastructure/db"
	"todo-app/internal/infrastructure/logger"
	"todo-app/internal/interface/handler"
	"todo-app/internal/interface/middleware"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func Init(e *echo.Echo, bunDB *bun.DB) {

	transaction := db.NewTransaction(bunDB)
	taskRepository := db.NewTaskRepository(bunDB)
	logger := logger.NewLogger()

	taskHandler := handler.NewTaskHandler(logger, transaction, taskRepository)

	loggerMiddleware := middleware.NewLoggerMiddleware(logger)

	e.Use(loggerMiddleware.Logger())

	e.POST("/tasks", taskHandler.CreateTask)
	e.GET("/tasks", taskHandler.GetTaskList)
	e.GET("/tasks/:id", taskHandler.GetTaskOne)
	e.PUT("/tasks/:id", taskHandler.UpdateTask)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)
}
