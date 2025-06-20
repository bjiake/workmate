package task

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	model "workmate/internal/domain/models/task"
	"workmate/internal/service"
)

type Repository struct {
	mu    sync.RWMutex
	tasks map[string]model.Task
}

func NewRepository() *Repository {
	repo := &Repository{
		tasks: make(map[string]model.Task),
	}
	return repo
}

func (this *Repository) Get(ctx context.Context, name string) (model.Task, error) {
	const op = "repo.task.Create"
	this.mu.RLock()
	defer this.mu.RUnlock()
	task, exists := this.tasks[name]
	if !exists {
		return model.Task{}, fmt.Errorf("%s: %w", op, service.ErrNotFound)
	}
	if task.Status != "Done" {
		task.Duration = time.Since(task.CreatedAt)
	}
	return task, nil
}

func (this *Repository) Create(ctx context.Context, task model.Task) error {
	const op = "repo.task.Create"
	this.mu.Lock()
	defer this.mu.Unlock()

	if _, exists := this.tasks[task.Name]; exists {
		return fmt.Errorf("%s: %w", op, errors.New("task already exists"))
	}

	this.tasks[task.Name] = task
	return nil
}

func (this *Repository) DoneTask(name string, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	task, exists := this.tasks[name]
	if !exists {
		return
	}
	task.Value = value
	task.Duration = time.Since(task.CreatedAt)
	task.Status = "Done"
	this.tasks[name] = task

	return
}

func (this *Repository) Delete(ctx context.Context, name string) {
	this.mu.Lock()
	defer this.mu.Unlock()
	delete(this.tasks, name)
}
