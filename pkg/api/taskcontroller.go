package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/api"
	"github.com/mole-squad/soq-api/pkg/auth"
	"github.com/mole-squad/soq-api/pkg/common"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
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

	dto := &api.CreateTaskRequestDto{}
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
		return
	}

	render.Render(w, r, task.AsDTO())
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

	dto := &api.UpdateTaskRequestDto{}
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
		return
	}

	render.Render(w, r, updatedTask.AsDTO())
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
		return
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
		return
	}

	respList := []render.Renderer{}
	for _, task := range userTasks {
		respList = append(respList, task.AsDTO())
	}

	render.RenderList(w, r, respList)
}
