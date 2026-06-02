package dto

type CreateRequest struct {
	ReleaseID string  `json:"release_id" binding:"required"`
	ErrorRate float64 `json:"error_rate" binding:"required"`
	LatencyMS float64 `json:"latency_ms" binding:"required"`
}