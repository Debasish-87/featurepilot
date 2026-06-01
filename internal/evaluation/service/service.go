package service

type Service struct {
	repo  Repository
	cache Cache
}

func New(
	repo Repository,
	cache Cache,
) *Service {

	return &Service{
		repo:  repo,
		cache: cache,
	}
}
