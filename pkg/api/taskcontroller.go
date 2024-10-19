package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/pkg/common"
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type TaskControllerParams struct {
	fx.In

	AuthService   interfaces.AuthService
	LoggerService interfaces.LoggerService
	TaskService   interfaces.TaskService
	Router        *chi.Mux
}

type TaskControllerResult struct {
	fx.Out

	TaskController TaskController
}

type TaskController struct {
	interfaces.ResourceController[*models.Task]

	logger      interfaces.LoggerService
	taskService interfaces.TaskService
}

func NewTaskController(params TaskControllerParams) (TaskControllerResult, error) {
	ctrl := TaskController{
		logger:      params.LoggerService,
		taskService: params.TaskService,
	}

	ctrl.ResourceController = generics.NewController[*models.Task](
		params.TaskService,
		params.LoggerService,
		params.AuthService,
		models.NewTaskFromCreateRequest,
		models.NewTaskFromUpdateRequest,
		generics.WithContextKey[*models.Task](taskContextkey),
		generics.WithDetailRoute[*models.Task]("PATCH", "/resolve", ctrl.ResolveTask),
	).(*generics.Controller[*models.Task])

	params.Router.Mount("/tasks", ctrl.ResourceController.GetRouter())

	return TaskControllerResult{TaskController: ctrl}, nil
}

func (ctrl *TaskController) ResolveTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	task, err := ctrl.ItemFromContext(ctx)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	updatedTask, err := ctrl.taskService.ResolveTask(ctx, task.ID)
	if err != nil {
		ctrl.logger.Error("failed to resolve task", "error", err)
		render.Render(w, r, common.ErrUnknown(err))

		return
	}

	render.Render(w, r, updatedTask.ToDTO())
}
