package dto

type ReleaseResponse struct {
	ID string `json:"id"`

	ProjectID string `json:"project_id"`

	Version string `json:"version"`

	Status string `json:"status"`
}