package service

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func (s *Service) Evaluate(
	ctx context.Context,
	environment string,
	featureKey string,
) (bool, error) {

	log.Println("EVALUATE CALLED")

	cacheKey := fmt.Sprintf(
		"feature:%s:%s",
		environment,
		featureKey,
	)

	log.Println("CACHE KEY =", cacheKey)

	enabled, err := s.cache.Get(
		ctx,
		cacheKey,
	)

	if err == nil {
		log.Println("CACHE HIT")
		return enabled, nil
	}

	if err == redis.Nil {
		log.Println("CACHE MISS")
	} else {
		log.Println("CACHE ERROR =", err)
	}

	enabled, err = s.repo.Evaluate(
		ctx,
		environment,
		featureKey,
	)

	if err != nil {
		return false, err
	}

	log.Println("DB RESULT =", enabled)

	err = s.cache.Set(
		ctx,
		cacheKey,
		enabled,
	)

	log.Println("SET RESULT =", err)

	return enabled, nil
}
