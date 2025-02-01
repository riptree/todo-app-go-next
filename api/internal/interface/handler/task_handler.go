package handler

import (
	"errors"
	"net/http"
	"strconv"
	"todo-app/internal/application/service"
	"todo-app/internal/application/usecase"
	"todo-app/internal/dto"
	"todo-app/internal/package/apperrors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

type taskHandler struct {
	logger      service.Logger
	taskUsecase *usecase.TaskUsecase
}

func NewTaskHandler(logger service.Logger, taskUsecase *usecase.TaskUsecase) *taskHandler {
	return &taskHandler{
		logger:      logger,
		taskUsecase: taskUsecase,
	}
}

func (h *taskHandler) CreateTask(c echo.Context) error {
	ctx := c.Request().Context()

	var params dto.CreateUpdateTaskRequest
	if err := c.Bind(&params); err != nil {
		h.logger.Errorf(ctx, "failed to bind request body: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "bad request",
		})
	}

	err := h.taskUsecase.CreateTask(ctx, params)
	if err != nil {
		h.logger.Errorf(ctx, "failed to CreateTask: %s", err.Error())

		if verr, ok := err.(validation.Errors); ok {
			return respondValidationError(verr)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": "failed to create task",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (h *taskHandler) UpdateTask(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf(ctx, "failed to parse id: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "invalid id",
		})
	}

	var params dto.CreateUpdateTaskRequest
	if err := c.Bind(&params); err != nil {
		h.logger.Errorf(ctx, "failed to bind request body: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "bad request",
		})
	}

	if err := h.taskUsecase.UpdateTask(ctx, id, params); err != nil {
		h.logger.Errorf(ctx, "failed to UpdateTask: %s", err.Error())

		if verr, ok := err.(validation.Errors); ok {
			return respondValidationError(verr)
		}

		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "failed to update task",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (h *taskHandler) DeleteTask(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf(ctx, "failed to parse id: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "invalid id",
		})
	}

	if err := h.taskUsecase.DeleteTask(ctx, id); err != nil {
		h.logger.Errorf(ctx, "failed to DeleteTask: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "failed to delete task",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (h *taskHandler) GetTaskList(c echo.Context) error {
	ctx := c.Request().Context()

	var params dto.GetTaskListRequest
	if err := c.Bind(&params); err != nil {
		h.logger.Errorf(ctx, "failed to bind query params: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": http.StatusText(http.StatusBadRequest),
		})
	}

	res, err := h.taskUsecase.GetTaskList(ctx, params)
	if err != nil {
		h.logger.Errorf(ctx, "failed to GetTaskList: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": http.StatusText(http.StatusInternalServerError),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *taskHandler) GetTaskOne(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf(ctx, "failed to parse id: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "invalid id",
		})
	}

	res, err := h.taskUsecase.GetTaskOne(ctx, id)
	if err != nil {
		h.logger.Errorf(ctx, "failed to GetTaskOne: %s", err.Error())
		if errors.Is(err, apperrors.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, map[string]any{
				"message": "task not found",
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": "failed to get task",
		})
	}

	return c.JSON(http.StatusOK, res)
}
