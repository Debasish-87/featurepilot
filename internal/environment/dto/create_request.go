package dto

type CreateRequest struct {
	ProjectID string `json:"project_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
}
