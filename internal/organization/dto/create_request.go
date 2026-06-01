package dto

type CreateOrganizationRequest struct {
	Name string `json:"name" binding:"required"`
}
