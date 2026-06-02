package dto

type CreateRequest struct {
	EnvironmentID string `json:"environment_id" binding:"required"`

	Key string `json:"key" binding:"required"`

	Name string `json:"name" binding:"required"`

	Description string `json:"description"`

	// 0-100
	RolloutPercentage int `json:"rollout_percentage"`
}
