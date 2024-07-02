package tasks

import (
	"net/http"

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
	fakeTasks := []Task{
		{Summary: "Test Task"},
		{Summary: "Test Task 2"},
		{Summary: "Test Task 3"},
		{Summary: "Test Task 4"},
	}

	render.RenderList(w, r, NewTaskListResponseDTO(fakeTasks))
}
