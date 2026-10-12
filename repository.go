package main

type HealthRepository interface {
	GetStatus() string
}

type InMemoryHelathRepository struct{}

func (r *InMemoryHelathRepository) GetStatus() string {
	return "healthy"
}
