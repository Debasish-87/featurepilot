package service

import releaseService "github.com/Debasish-87/featurepilot/internal/release/service"

type Service struct {
	repo Repository

	releaseSvc *releaseService.Service
}

func New(
	repo Repository,
	releaseSvc *releaseService.Service,
) *Service {

	return &Service{
		repo:       repo,
		releaseSvc: releaseSvc,
	}
}
