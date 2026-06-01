package service

import "context"

type EvaluationCache interface {
	Delete(
		ctx context.Context,
		key string,
	) error
}
