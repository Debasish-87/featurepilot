package service

type Service struct {
	repo      Repository
	envRepo   EnvironmentRepository
	evalCache EvaluationCache
}

func New(
	repo Repository,
	envRepo EnvironmentRepository,
	evalCache EvaluationCache,
) *Service {

	return &Service{
		repo:      repo,
		envRepo:   envRepo,
		evalCache: evalCache,
	}
}
