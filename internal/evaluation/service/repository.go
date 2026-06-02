package service

import "context"

type EvaluationResult struct {
	Enabled           bool
	RolloutPercentage int
}

type Repository interface {
	Evaluate(
		ctx context.Context,
		environment string,
		featureKey string,
	) (*EvaluationResult, error)
}
