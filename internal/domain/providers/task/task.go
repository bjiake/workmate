package Task

import (
	"github.com/google/wire"
	"log/slog"
	"sync"
	taskHdl "workmate/internal/api/handler/task"
	"workmate/internal/domain/interfaces"
	taskRepo "workmate/internal/repo/task"
	taskSvc "workmate/internal/service/task"
)

var (
	hdl     *taskHdl.TaskHandler
	hdlOnce sync.Once

	svc     *taskSvc.Service
	svcOnce sync.Once

	repo     *taskRepo.Repository
	repoOnce sync.Once
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	ProvideTaskHandler,
	ProvideTaskService,
	ProvideTaskRepository,

	wire.Bind(new(interfaces.TaskHandler), new(*taskHdl.TaskHandler)),
	wire.Bind(new(interfaces.TaskService), new(*taskSvc.Service)),
	wire.Bind(new(interfaces.TaskRepository), new(*taskRepo.Repository)),
)

func ProvideTaskHandler(svc interfaces.TaskService, log *slog.Logger) *taskHdl.TaskHandler {
	hdlOnce.Do(func() {
		hdl = &taskHdl.TaskHandler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideTaskService(TaskRepo interfaces.TaskRepository, someService interfaces.SomethingService) *taskSvc.Service {
	svcOnce.Do(func() {
		svc = &taskSvc.Service{
			TaskRepo:         TaskRepo,
			SomethingService: someService,
		}
	})

	return svc
}

func ProvideTaskRepository() *taskRepo.Repository {
	repoOnce.Do(func() {
		repo = taskRepo.NewRepository()
	})

	return repo
}
