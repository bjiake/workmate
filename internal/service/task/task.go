package task

import (
	"context"
	"fmt"
	"time"
	"workmate/internal/domain/interfaces"
	model "workmate/internal/domain/models/task"
)

type Service struct {
	TaskRepo         interfaces.TaskRepository
	SomethingService interfaces.SomethingService
}

func (this *Service) GetTaskByName(ctx context.Context, name string) (model.Task, error) {
	const op = "service.task.GetTaskByName"

	result, err := this.TaskRepo.Get(ctx, name)
	if err != nil {
		return model.Task{}, fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}

func (this *Service) DeleteTaskByName(ctx context.Context, name string) {
	const op = "service.task.DeleteTaskByName"

	this.TaskRepo.Delete(ctx, name)
	return
}

func (this *Service) CreateTask(ctx context.Context, newTask model.Create) (model.Task, error) {
	const op = "service.task.Create"
	currentTask := model.Task{
		Name:      newTask.Name,
		Status:    "Working",
		Value:     0,
		Duration:  time.Duration(0),
		CreatedAt: time.Now(),
	}

	if err := this.TaskRepo.Create(ctx, currentTask); err != nil {
		return model.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	go func() {
		result := this.SomethingService.Process(newTask.Value)
		this.TaskRepo.DoneTask(currentTask.Name, result)
	}()

	return currentTask, nil
}
