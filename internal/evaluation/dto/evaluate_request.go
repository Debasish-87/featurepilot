package dto

type EvaluateRequest struct {
	Environment string `json:"environment" binding:"required"`

	FeatureKey string `json:"feature_key" binding:"required"`

	UserID string `json:"user_id" binding:"required"`
}
