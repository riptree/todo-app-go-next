package dto

type CreateUpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

type GetTaskListRequest struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type GetTaskListResponse struct {
	Tasks []GetTaskResponse `json:"tasks"`
}

type GetTaskResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
