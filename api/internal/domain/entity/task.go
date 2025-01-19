package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks"`

	ID        int       `bun:",pk,autoincrement"`
	Title     string    `bun:"title"`
	Content   string    `bun:"content"`
	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`
}

func NewTask(title string, content string) (Task, error) {
	return Task{
		Title:   title,
		Content: content,
	}, nil
}

func (t *Task) UpdateTask(title string, content string) {
	t.Title = title
	t.Content = content
}
