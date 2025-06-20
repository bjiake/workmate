package interfaces

import (
	"context"
	"net/http"
	model "workmate/internal/domain/models/task"
)

type TaskRepository interface {
	Get(ctx context.Context, name string) (model.Task, error)
	Create(ctx context.Context, task model.Task) error
	DoneTask(name string, value int)
	Delete(ctx context.Context, name string)
}

type TaskService interface {
	GetTaskByName(ctx context.Context, name string) (model.Task, error)
	DeleteTaskByName(ctx context.Context, idStr string)
	CreateTask(ctx context.Context, task model.Create) (model.Task, error)
}

type TaskHandler interface {
	GetTaskByName(w http.ResponseWriter, r *http.Request)
	DeleteTaskByName(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
}
