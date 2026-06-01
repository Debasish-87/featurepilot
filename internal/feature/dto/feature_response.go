package dto

type FeatureResponse struct {
	ID string `json:"id"`

	EnvironmentID string `json:"environment_id"`

	Key string `json:"key"`

	Name string `json:"name"`

	Description string `json:"description"`

	Enabled bool `json:"enabled"`
}
