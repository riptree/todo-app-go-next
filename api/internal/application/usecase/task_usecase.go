package usecase

import (
	"context"
	"task-management/internal/application/service"
	"task-management/internal/domain/entity"
	"task-management/internal/domain/repository"
	"task-management/internal/dto"
	"time"

	"github.com/samber/lo"
)

type TaskUsecase interface {
	CreateTask(ctx context.Context, req dto.CreateUpdateTaskRequest) error
	UpdateTask(ctx context.Context, id int, req dto.CreateUpdateTaskRequest) error
	DeleteTask(ctx context.Context, id int) error
	GetTaskList(ctx context.Context, req dto.GetTaskListRequest) (dto.GetTaskListResponse, error)
	GetTaskOne(ctx context.Context, id int) (dto.GetTaskResponse, error)
}

type taskUsecase struct {
	transaction    service.Transaction
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(
	transaction service.Transaction,
	taskRepository repository.TaskRepository,
) TaskUsecase {
	return &taskUsecase{
		transaction:    transaction,
		taskRepository: taskRepository,
	}
}

func (u *taskUsecase) CreateTask(ctx context.Context, req dto.CreateUpdateTaskRequest) error {

	err := u.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		task, err := entity.NewTask(req.Title, req.Content)
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

func (u *taskUsecase) UpdateTask(ctx context.Context, id int, req dto.CreateUpdateTaskRequest) error {

	err := u.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		task, err := u.taskRepository.GetTaskOne(ctx, id)
		if err != nil {
			return err
		}

		task.UpdateTask(req.Title, req.Content)

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

func (u *taskUsecase) DeleteTask(ctx context.Context, id int) error {

	err := u.taskRepository.DeleteTask(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *taskUsecase) GetTaskList(ctx context.Context, req dto.GetTaskListRequest) (dto.GetTaskListResponse, error) {

	resTasks := dto.GetTaskListResponse{
		Tasks: []dto.GetTaskResponse{},
	}

	// limit default is 100
	l := req.Limit
	if l == 0 {
		l = 100
	}

	tasks, err := u.taskRepository.GetTaskList(ctx, l, req.Offset)
	if err != nil {
		return resTasks, err
	}

	resTasks = dto.GetTaskListResponse{
		Tasks: lo.Map(tasks, func(t entity.Task, _ int) dto.GetTaskResponse {
			return dto.GetTaskResponse{
				ID:        t.ID,
				Title:     t.Title,
				Content:   t.Content,
				CreatedAt: t.CreatedAt.Format(time.RFC3339),
				UpdatedAt: t.UpdatedAt.Format(time.RFC3339),
			}
		}),
	}

	return resTasks, nil
}

func (u *taskUsecase) GetTaskOne(ctx context.Context, id int) (dto.GetTaskResponse, error) {
	var resTask dto.GetTaskResponse

	task, err := u.taskRepository.GetTaskOne(ctx, id)
	if err != nil {
		return resTask, err
	}

	resTask = dto.GetTaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		CreatedAt: task.CreatedAt.Format(time.RFC3339),
		UpdatedAt: task.UpdatedAt.Format(time.RFC3339),
	}

	return resTask, nil
}
