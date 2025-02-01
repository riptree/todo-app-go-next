package router

import (
	"todo-app/internal/infrastructure/db"
	"todo-app/internal/infrastructure/logger"
	"todo-app/internal/interface/controller"
	"todo-app/internal/interface/middleware"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func Init(e *echo.Echo, bunDB *bun.DB) {

	transaction := db.NewTransaction(bunDB)
	taskRepository := db.NewTaskRepository(bunDB)
	logger := logger.NewLogger()

	taskController := controller.NewTaskController(logger, transaction, taskRepository)

	loggerMiddleware := middleware.NewLoggerMiddleware(logger)

	e.Use(loggerMiddleware.Logger())

	e.POST("/tasks", taskController.CreateTask)
	e.GET("/tasks", taskController.GetTaskList)
	e.GET("/tasks/:id", taskController.GetTaskOne)
	e.PUT("/tasks/:id", taskController.UpdateTask)
	e.DELETE("/tasks/:id", taskController.DeleteTask)
}
