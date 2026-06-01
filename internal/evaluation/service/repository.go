package service

import "context"

type Repository interface {
	Evaluate(
		ctx context.Context,
		environment string,
		featureKey string,
	) (bool, error)
}
