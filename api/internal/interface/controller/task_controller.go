package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"todo-app/internal/application/service"
	"todo-app/internal/domain/entity"
	"todo-app/internal/dto"
	"todo-app/internal/infrastructure/db"
	"todo-app/internal/package/apperrors"
	"todo-app/internal/package/util"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type taskController struct {
	logger         service.Logger
	transaction    service.Transaction
	taskRepository *db.TaskRepository
}

func NewTaskController(
	logger service.Logger,
	transaction service.Transaction,
	taskRepository *db.TaskRepository,
) *taskController {
	return &taskController{
		logger:         logger,
		transaction:    transaction,
		taskRepository: taskRepository,
	}
}

func (h *taskController) CreateTask(c echo.Context) error {
	ctx := c.Request().Context()

	var params dto.CreateUpdateTaskRequest
	if err := c.Bind(&params); err != nil {
		h.logger.Errorf(ctx, "failed to bind request body: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "bad request",
		})
	}

	if err := params.Validate(); err != nil {
		h.logger.Errorf(ctx, "failed to CreateTask: %s", err.Error())

		if verr, ok := err.(validation.Errors); ok {
			return respondValidationError(verr)
		}

		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "failed to create task",
		})
	}

	err := h.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		dueDate, err := util.ParseDate(params.DueDate)
		if err != nil {
			return err
		}

		task, err := entity.NewTask(params.Title, params.Description, dueDate)
		if err != nil {
			return err
		}

		_, err = h.taskRepository.CreateTask(ctx, &task)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		h.logger.Errorf(ctx, "failed to CreateTask: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": "failed to create task",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (h *taskController) UpdateTask(c echo.Context) error {
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

	if err := params.Validate(); err != nil {
		h.logger.Errorf(ctx, "validation error: %s", err.Error())

		if verr, ok := err.(validation.Errors); ok {
			return respondValidationError(verr)
		}

		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "failed to update task",
		})
	}

	err = h.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		task, err := h.taskRepository.GetTaskOne(ctx, id)
		if err != nil {
			return err
		}

		dueDate, err := util.ParseDate(params.DueDate)
		if err != nil {
			return err
		}

		task.UpdateTask(params.Title, params.Description, dueDate)

		err = h.taskRepository.UpdateTask(ctx, &task)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		h.logger.Errorf(ctx, "failed to CreateTask: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": "failed to update task",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (h *taskController) DeleteTask(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf(ctx, "failed to parse id: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "invalid id",
		})
	}

	err = h.taskRepository.DeleteTask(ctx, id)
	if err != nil {
		h.logger.Errorf(ctx, "failed to DeleteTask: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "failed to delete task",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (h *taskController) GetTaskList(c echo.Context) error {
	ctx := c.Request().Context()

	var params dto.GetTaskListRequest
	if err := c.Bind(&params); err != nil {
		h.logger.Errorf(ctx, "failed to bind query params: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": http.StatusText(http.StatusBadRequest),
		})
	}

	res := dto.GetTaskListResponse{
		Tasks: []dto.GetTaskResponse{},
	}

	// limit default is 100
	l := params.Limit
	if l == 0 {
		l = 100
	}

	tasks, err := h.taskRepository.GetTaskList(ctx, l, params.Offset)
	if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
		h.logger.Errorf(ctx, "failed to GetTaskList: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": http.StatusText(http.StatusInternalServerError),
		})
	}

	res = dto.GetTaskListResponse{
		Tasks: lo.Map(tasks, func(t entity.Task, _ int) dto.GetTaskResponse {
			return dto.GetTaskResponse{
				ID:          t.ID,
				Title:       t.Title,
				Description: t.Description,
				DueDate:     t.GetDueDateString(),
				CreatedAt:   t.GetCreatedAtString(),
				UpdatedAt:   t.GetUpdatedAtString(),
			}
		}),
	}

	return c.JSON(http.StatusOK, res)
}

func (h *taskController) GetTaskOne(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf(ctx, "failed to parse id: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"message": "invalid id",
		})
	}

	var res dto.GetTaskResponse

	task, err := h.taskRepository.GetTaskOne(ctx, id)
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

	res = dto.GetTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.GetDueDateString(),
		CreatedAt:   task.GetCreatedAtString(),
		UpdatedAt:   task.GetUpdatedAtString(),
	}

	return c.JSON(http.StatusOK, res)
}
