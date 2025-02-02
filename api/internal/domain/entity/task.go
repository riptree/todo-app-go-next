package entity

import (
	"fmt"
	"time"
	"todo-app/internal/package/util"

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

func (t *Task) PatchUpdateTask(m map[string]any) error {
	if title, ok := m["title"]; ok {
		var valid bool
		t.Title, valid = title.(string)
		if !valid {
			return fmt.Errorf("invalid title")
		}
	}

	if description, ok := m["description"]; ok {
		if description == nil {
			t.Description = nil
		} else {
			descriptionString, valid := description.(string)
			if !valid {
				return fmt.Errorf("invalid description")
			}
			t.Description = &descriptionString
		}
	}

	if dueDate, ok := m["due_date"]; ok {
		if dueDate == nil {
			t.DueDate = nil
		} else {
			dueDateString, valid := dueDate.(string)
			if !valid {
				return fmt.Errorf("invalid due_date")
			}

			dueDate, err := util.ParseDate(&dueDateString)
			if err != nil {
				return err
			}
			t.DueDate = dueDate
		}
	}

	return nil
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
