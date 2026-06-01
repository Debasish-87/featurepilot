package service

import "context"

type Cache interface {
	Get(
		ctx context.Context,
		key string,
	) (bool, error)

	Set(
		ctx context.Context,
		key string,
		enabled bool,
	) error

	Delete(
		ctx context.Context,
		key string,
	) error
}
