package repository

import (
	"context"
	"task-management/internal/domain/entity"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, data *entity.Task) (int, error)
	UpdateTask(ctx context.Context, data *entity.Task) error
	DeleteTask(ctx context.Context, taskID int) error
	GetTaskList(ctx context.Context, limit int, offset int) ([]entity.Task, error)
	GetTaskOne(ctx context.Context, taskID int) (entity.Task, error)
}
