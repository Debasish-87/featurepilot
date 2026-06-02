package dto

type CreateRequest struct {
	ProjectID string `json:"project_id" binding:"required"`
	Version   string `json:"version" binding:"required"`
}