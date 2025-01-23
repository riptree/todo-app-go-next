package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks"`

	ID          int        `bun:",pk,autoincrement"`
	Title       string     `bun:"title"`
	Description *string    `bun:"description"`
	DueDate     *time.Time `bun:"due_date"`
	CreatedAt   time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time  `bun:"updated_at,notnull,default:current_timestamp"`
}

func NewTask(title string, description *string, dueDate *time.Time) (Task, error) {
	return Task{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}, nil
}

func (t *Task) UpdateTask(title string, description *string, dueDate *time.Time) {
	t.Title = title
	t.Description = description
	t.DueDate = dueDate
}

func (t *Task) IsDue() bool {
	return t.DueDate.Before(time.Now())
}

func (t *Task) GetDueDateString() *string {
	if t.DueDate == nil {
		return nil
	}
	dueDateString := t.DueDate.Format("2006-01-02")
	return &dueDateString
}

func (t *Task) GetCreatedAtString() string {
	return t.CreatedAt.Format(time.RFC3339)
}

func (t *Task) GetUpdatedAtString() string {
	return t.UpdatedAt.Format(time.RFC3339)
}
