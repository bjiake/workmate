package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"workmate/internal/domain/interfaces"
	_ "workmate/internal/domain/models/response"
	model "workmate/internal/domain/models/task"
	"workmate/internal/service"
	"workmate/pkg/logger/slogError"
	responseApi "workmate/pkg/response"
)

type TaskHandler struct {
	Svc interfaces.TaskService
	Log *slog.Logger
}

func (h *TaskHandler) NewTaskHandler(r chi.Router) {
	r.Route("/task", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/{name}", h.GetTaskByName)
			r.Post("/", h.CreateTask)
			r.Delete("/{name}", h.DeleteTaskByName)
		})
	})
}

// GetTaskByName godoc
// @Summary Получить задачу по имени
// @Description Получает задачу по имени
// @Tags tasks
// @Accept json
// @Produce json
// @Param name path string true "Имя задачи"
// @Success 200 {object} task.Task
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Router /task/{name} [get]
func (h *TaskHandler) GetTaskByName(w http.ResponseWriter, r *http.Request) {
	const op = "handler.task.GetTaskByName"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var name = chi.URLParam(r, "name")

	result, err := h.Svc.GetTaskByName(context.Background(), name)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			responseApi.WriteError(w, r, http.StatusNotFound, slogError.Err(err))
			return
		}
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, result)

}

// CreateTask godoc
// @Summary Создать новую задачу
// @Description Создает новую задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.Create true "Создание задачи"
// @Success 201 {object} task.Task
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /task [post]
func (h *TaskHandler) DeleteTaskByName(w http.ResponseWriter, r *http.Request) {
	const op = "handler.task.DeleteTaskByName"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	var name = chi.URLParam(r, "name")

	h.Svc.DeleteTaskByName(context.Background(), name)
	responseApi.WriteJson(w, r, http.StatusNoContent, nil)
}

// DeleteTaskByName godoc
// @Summary Удалить задачу по имени
// @Description Удаляет задачу по имени
// @Tags tasks
// @Accept json
// @Produce json
// @Param name path string true "Имя задачи"
// @Success 204
// @Failure 404 {object} response.Error
// @Router /task/{name} [delete]
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	const op = "handler.task.CreateTask"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var newTask model.Create
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	result, err := h.Svc.CreateTask(r.Context(), newTask)
	if err != nil {
		h.Log.Error("service failed to create task profile picture", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, result)
}
