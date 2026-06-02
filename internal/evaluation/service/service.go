package service

type Service struct {
	repo Repository

	cache Cache

	eventRepo EventRepository
}

func New(
	repo Repository,
	cache Cache,
	eventRepo EventRepository,
) *Service {

	return &Service{
		repo:      repo,
		cache:     cache,
		eventRepo: eventRepo,
	}
}
