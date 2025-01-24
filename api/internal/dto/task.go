package dto

import (
	"todo-app/internal/package/global"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateUpdateTaskRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	DueDate     *string `json:"due_date"`
}

func (t CreateUpdateTaskRequest) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Title, validation.Required, validation.Length(1, 255)),
		validation.Field(&t.DueDate, validation.Date(global.DateFormat)),
	)
}

type GetTaskListRequest struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type GetTaskListResponse struct {
	Tasks []GetTaskResponse `json:"tasks"`
}

type GetTaskResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	DueDate     *string `json:"due_date"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
