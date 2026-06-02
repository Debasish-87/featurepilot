package service

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (s *Service) Enable(
	ctx context.Context,
	id uuid.UUID,
) error {

	log.Println("ENABLE CALLED:", id)

	feature, err := s.repo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}

	if err := s.repo.Enable(
		ctx,
		id,
	); err != nil {
		return err
	}

	environment, err := s.envRepo.GetByID(
		ctx,
		feature.EnvironmentID,
	)
	if err != nil {
		return err
	}

	log.Printf("AUDIT OBJECT: %+v\n", s.audit)

	// Audit Event
	if s.audit != nil {

		log.Println("ABOUT TO WRITE AUDIT")

		s.audit.Log(
			ctx,
			"FEATURE_ENABLED",
			"feature",
			id,
			fmt.Sprintf(
				`{"key":"%s","environment":"%s"}`,
				feature.Key,
				environment.Name,
			),
		)

		log.Println("AUDIT CALL FINISHED")
	} else {
		log.Println("AUDIT IS NIL")
	}

	cacheKey := fmt.Sprintf(
		"feature:%s:%s",
		environment.Name,
		feature.Key,
	)

	_ = s.evalCache.Delete(
		ctx,
		cacheKey,
	)

	return nil
}
