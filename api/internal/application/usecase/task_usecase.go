package usecase

import (
	"context"
	"errors"
	"todo-app/internal/application/service"
	"todo-app/internal/domain/entity"
	"todo-app/internal/dto"
	"todo-app/internal/infrastructure/db"
	"todo-app/internal/package/apperrors"
	"todo-app/internal/package/util"

	"github.com/samber/lo"
)

type TaskUsecase struct {
	transaction    service.Transaction
	taskRepository *db.TaskRepository
}

func NewTaskUsecase(
	transaction service.Transaction,
	taskRepository *db.TaskRepository,
) *TaskUsecase {
	return &TaskUsecase{
		transaction:    transaction,
		taskRepository: taskRepository,
	}
}

func (u *TaskUsecase) CreateTask(ctx context.Context, req dto.CreateUpdateTaskRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	err := u.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		dueDate, err := util.ParseDate(req.DueDate)
		if err != nil {
			return err
		}

		task, err := entity.NewTask(req.Title, req.Description, dueDate)
		if err != nil {
			return err
		}

		_, err = u.taskRepository.CreateTask(ctx, &task)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *TaskUsecase) UpdateTask(ctx context.Context, id int, req dto.CreateUpdateTaskRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	err := u.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		task, err := u.taskRepository.GetTaskOne(ctx, id)
		if err != nil {
			return err
		}

		dueDate, err := util.ParseDate(req.DueDate)
		if err != nil {
			return err
		}

		task.UpdateTask(req.Title, req.Description, dueDate)

		err = u.taskRepository.UpdateTask(ctx, &task)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *TaskUsecase) DeleteTask(ctx context.Context, id int) error {

	err := u.taskRepository.DeleteTask(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *TaskUsecase) GetTaskList(ctx context.Context, req dto.GetTaskListRequest) (dto.GetTaskListResponse, error) {

	resTasks := dto.GetTaskListResponse{
		Tasks: []dto.GetTaskResponse{},
	}

	// limit default is 100
	l := req.Limit
	if l == 0 {
		l = 100
	}

	tasks, err := u.taskRepository.GetTaskList(ctx, l, req.Offset)
	if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
		return resTasks, err
	}

	resTasks = dto.GetTaskListResponse{
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

	return resTasks, nil
}

func (u *TaskUsecase) GetTaskOne(ctx context.Context, id int) (dto.GetTaskResponse, error) {
	var resTask dto.GetTaskResponse

	task, err := u.taskRepository.GetTaskOne(ctx, id)
	if err != nil {
		return resTask, err
	}

	resTask = dto.GetTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.GetDueDateString(),
		CreatedAt:   task.GetCreatedAtString(),
		UpdatedAt:   task.GetUpdatedAtString(),
	}

	return resTask, nil
}
