package tasks

import (
	"net/http"

	"github.com/burkel24/task-app/pkg/auth"
	"github.com/burkel24/task-app/pkg/common"
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/fx"
)

type TaskControllerParams struct {
	fx.In

	TaskService interfaces.TaskService
	Router      *chi.Mux
}

type TaskControllerResult struct {
	fx.Out

	TaskController interface{}
}

type TaskController struct {
	taskService interfaces.TaskService
}

func NewTaskController(params TaskControllerParams) (TaskControllerResult, error) {
	ctrl := TaskController{taskService: params.TaskService}

	taskRouter := chi.NewRouter()
	taskRouter.Get("/", ctrl.ListTasks)

	params.Router.Mount("/tasks", taskRouter)

	return TaskControllerResult{TaskController: ctrl}, nil
}

func (ctrl *TaskController) ListTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	tasks, err := ctrl.taskService.ListUserTasks(ctx, user)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
	}

	render.RenderList(w, r, NewTaskListResponseDTO(tasks))
}
