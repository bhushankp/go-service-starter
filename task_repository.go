package main

import (
	"errors"
	"sync"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository interface {
	Create(task Task) Task
	GetAll() []Task
	GetByID(id int) (Task, error)
	Update(id int, task Task) (Task, error)
	Delete(id int) error
}

type InMemoryTaskRepository struct {
	mu     sync.RWMutex
	tasks  []Task
	nextID int
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

func (r *InMemoryTaskRepository) Create(task Task) Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = r.nextID
	r.nextID++

	r.tasks = append(r.tasks, task)

	return task
}

func (r *InMemoryTaskRepository) GetAll() []Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]Task, len(r.tasks))
	copy(result, r.tasks)

	return result
}

func (r *InMemoryTaskRepository) GetByID(id int) (Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, ErrTaskNotFound
}

func (r *InMemoryTaskRepository) Update(id int, updated Task) (Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, task := range r.tasks {
		if task.ID == id {
			updated.ID = id
			if updated.Title == "" {
				return Task{}, errors.New("title is required")
			}
			if updated.Status == "" {
				updated.Status = task.Status
			}
			r.tasks[i] = updated
			return updated, nil
		}
	}
	return Task{}, ErrTaskNotFound
}

func (r *InMemoryTaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return ErrTaskNotFound
}
