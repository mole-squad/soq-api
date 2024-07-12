package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/burkel24/task-app/pkg/auth"
	"github.com/burkel24/task-app/pkg/common"
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
	"github.com/burkel24/task-app/pkg/tasks"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type TaskControllerParams struct {
	fx.In

	AuthService interfaces.AuthService
	TaskService interfaces.TaskService
	Router      *chi.Mux
}

type TaskControllerResult struct {
	fx.Out

	TaskController TaskController
}

type TaskController struct {
	taskService interfaces.TaskService
}

func NewTaskController(params TaskControllerParams) (TaskControllerResult, error) {
	ctrl := TaskController{taskService: params.TaskService}

	taskRouter := chi.NewRouter()
	taskRouter.Use(params.AuthService.AuthRequired())

	taskRouter.Get("/", ctrl.ListTasks)
	taskRouter.Post("/", ctrl.CreateTask)
	taskRouter.Patch("/{taskID}", ctrl.UpdateTask)
	taskRouter.Delete("/{taskID}", ctrl.DeleteTask)

	params.Router.Mount("/tasks", taskRouter)

	return TaskControllerResult{TaskController: ctrl}, nil
}

func (ctrl *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	dto := &tasks.CreateTaskRequestDto{}
	if err = render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	// TODO validate user owns focus area
	newTask := models.Task{
		Summary:     dto.Summary,
		Notes:       dto.Notes,
		FocusAreaID: dto.FocusAreaID,
	}

	task, err := ctrl.taskService.CreateUserTask(ctx, user, &newTask)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
	}

	resp := tasks.NewTaskDTO(task)
	render.Render(w, r, resp)
}

func (ctrl *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	taskId := chi.URLParam(r, "taskID")
	taskIdInt, err := strconv.Atoi(taskId)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("failed to parse taskID: %w", err)))
	}

	dto := &tasks.UpdateTaskRequestDto{}
	if err = render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	// TODO validate user owns focus area
	task := models.Task{
		Model:       gorm.Model{ID: uint(taskIdInt)},
		Summary:     dto.Summary,
		Notes:       dto.Notes,
		FocusAreaID: dto.FocusAreaID,
	}

	updatedTask, err := ctrl.taskService.UpdateUserTask(ctx, &task)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
	}

	resp := tasks.NewTaskDTO(updatedTask)
	render.Render(w, r, resp)
}

func (ctrl *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	taskId := chi.URLParam(r, "taskID")
	taskIdInt, err := strconv.Atoi(taskId)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("failed to parse taskID: %w", err)))
	}

	err = ctrl.taskService.DeleteUserTask(ctx, uint(taskIdInt))
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
	}

	render.NoContent(w, r)
}

func (ctrl *TaskController) ListTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	userTasks, err := ctrl.taskService.ListOpenUserTasks(ctx, user.ID)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
	}

	render.RenderList(w, r, tasks.NewTaskListResponseDTO(userTasks))
}
