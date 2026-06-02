package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	eventDomain "github.com/Debasish-87/featurepilot/internal/evaluation_event/domain"
	"github.com/Debasish-87/featurepilot/internal/metrics"
	"github.com/Debasish-87/featurepilot/pkg/utils"
)

func (s *Service) Evaluate(
	ctx context.Context,
	environment string,
	featureKey string,
	userID string,
) (bool, error) {

	start := time.Now()

	metrics.EvaluationTotal.Inc()

	log.Println("EVALUATE CALLED")

	cacheKey := fmt.Sprintf(
		"feature:%s:%s:%s",
		environment,
		featureKey,
		userID,
	)

	log.Println("CACHE KEY =", cacheKey)

	enabled, err := s.cache.Get(
		ctx,
		cacheKey,
	)

	if err == nil {

		metrics.CacheHitTotal.Inc()

		log.Println("CACHE HIT")

		metrics.EvaluationDuration.Observe(
			time.Since(start).Seconds(),
		)

		return enabled, nil
	}

	if err == redis.Nil {

		metrics.CacheMissTotal.Inc()

		log.Println("CACHE MISS")

	} else {

		log.Println("CACHE ERROR =", err)
	}

	result, err := s.repo.Evaluate(
		ctx,
		environment,
		featureKey,
	)

	if err != nil {

		metrics.EvaluationErrorsTotal.Inc()

		return false, err
	}

	if !result.Enabled {

		log.Println("FEATURE DISABLED")

		return false, nil
	}

	bucket := utils.RolloutBucket(
		userID,
	)

	enabled = bucket < result.RolloutPercentage

	log.Printf(
		"ROLLOUT bucket=%d rollout=%d enabled=%v",
		bucket,
		result.RolloutPercentage,
		enabled,
	)

	err = s.cache.Set(
		ctx,
		cacheKey,
		enabled,
	)

	if err != nil {
		log.Println("SET RESULT =", err)
	} else {
		log.Println("SET RESULT = SUCCESS")
	}

	metrics.EvaluationDuration.Observe(
		time.Since(start).Seconds(),
	)

	event := &eventDomain.EvaluationEvent{
		ID: uuid.New(),

		Environment: environment,
		FeatureKey:  featureKey,
		UserID:      userID,

		Enabled: enabled,

		CreatedAt: time.Now().UTC(),
	}

	go func() {
		err := s.eventRepo.Create(
			context.Background(),
			event,
		)

		if err != nil {
			log.Println(
				"EVENT INSERT FAILED:",
				err,
			)
		}
	}()

	return enabled, nil
}
