package dto

type CreateProjectRequest struct {
	OrganizationID string `json:"organization_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
}
