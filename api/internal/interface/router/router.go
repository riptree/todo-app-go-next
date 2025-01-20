package router

import (
	"task-management/internal/application/usecase"
	"task-management/internal/infrastructure/db"
	"task-management/internal/infrastructure/logger"
	"task-management/internal/interface/handler"
	"task-management/internal/interface/middleware"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func Init(e *echo.Echo, bunDB *bun.DB) {

	transaction := db.NewTransaction(bunDB)
	taskRepository := db.NewTaskRepository(bunDB)
	logger := logger.NewLogger()

	taskUsecase := usecase.NewTaskUsecase(transaction, taskRepository)

	taskHandler := handler.NewTaskHandler(logger, taskUsecase)

	loggerMiddleware := middleware.NewLoggerMiddleware(logger)

	e.Use(loggerMiddleware.Logger())

	e.POST("/tasks", taskHandler.CreateTask)
	e.GET("/tasks", taskHandler.GetTaskList)
	e.GET("/tasks/:id", taskHandler.GetTaskOne)
	e.PUT("/tasks/:id", taskHandler.UpdateTask)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)
}
