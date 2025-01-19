package dto

type CreateUpdateTaskRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetTaskListRequest struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type GetTaskListResponse struct {
	Tasks []GetTaskResponse `json:"tasks"`
}

type GetTaskResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
