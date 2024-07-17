package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/pkg/common"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"github.com/mole-squad/soq-api/pkg/rest"
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
	ctrl        *rest.Controller[*models.Task]
	taskService interfaces.TaskService
}

func NewTaskController(params TaskControllerParams) (TaskControllerResult, error) {
	taskCtrl := TaskController{taskService: params.TaskService}

	taskCtrl.ctrl = rest.NewController[*models.Task](
		params.TaskService,
		params.LoggerService,
		params.AuthService,
		models.NewTaskFromCreateRequest,
		models.NewTaskFromUpdateRequest,
		rest.WithContextKey[*models.Task](taskContextkey),
		rest.WithDetailRoute[*models.Task]("POST", "/resolve", taskCtrl.ResolveTask),
	)

	params.Router.Mount("/tasks", taskCtrl.ctrl.Router)

	return TaskControllerResult{TaskController: taskCtrl}, nil
}

func (ctrl *TaskController) ResolveTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	task, err := ctrl.ctrl.ItemFromContext(ctx)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
	}

	updatedTask, err := ctrl.taskService.ResolveUserTask(ctx, task.UserID, task.ID)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Render(w, r, updatedTask.ToDTO())
}
