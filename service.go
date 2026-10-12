package main

type HealthService struct {
	repo HealthRepository
}

func NewHealthService(repo HealthRepository) *HealthService {
	return &HealthService{repo: repo}
}

func (s *HealthService) Check() string {
	return s.repo.GetStatus()
}
