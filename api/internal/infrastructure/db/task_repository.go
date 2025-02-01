package db

import (
	"context"
	"database/sql"
	"errors"
	"todo-app/internal/domain/entity"
	"todo-app/internal/package/apperrors"

	"github.com/uptrace/bun"
)

type TaskRepository struct {
	conn *bun.DB
}

func NewTaskRepository(conn *bun.DB) *TaskRepository {
	return &TaskRepository{
		conn: conn,
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *entity.Task) (int, error) {
	tx := GetTxOrDB(ctx, r.conn)
	_, err := tx.NewInsert().Model(task).Exec(ctx)
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	tx := GetTxOrDB(ctx, r.conn)
	_, err := tx.NewUpdate().Model(task).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, taskID int) error {
	tx := GetTxOrDB(ctx, r.conn)
	_, err := tx.NewDelete().Model(&entity.Task{}).Where("id = ?", taskID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) GetTaskList(ctx context.Context, limit int, offset int) ([]entity.Task, error) {
	tasks := make([]entity.Task, 0, limit)

	tx := GetTxOrDB(ctx, r.conn)
	if err := tx.NewSelect().Model(&tasks).Limit(limit).Offset(offset).Scan(ctx); err != nil {
		return []entity.Task{}, err
	}

	if len(tasks) == 0 {
		return []entity.Task{}, apperrors.ErrNotFound
	}

	return tasks, nil
}

func (r *TaskRepository) GetTaskOne(ctx context.Context, taskID int) (entity.Task, error) {
	var task entity.Task

	tx := GetTxOrDB(ctx, r.conn)
	if err := tx.NewSelect().Model(&task).Where("id = ?", taskID).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Task{}, apperrors.ErrNotFound
		}
		return entity.Task{}, err
	}

	return task, nil
}
