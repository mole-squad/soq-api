package api

import (
	"context"
	"errors"
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

	taskRouter.Route("/{taskID}", func(r chi.Router) {
		r.Use(ctrl.taskContextMiddleware)

		r.Patch("/", ctrl.UpdateTask)
		r.Delete("/", ctrl.DeleteTask)
		r.Post("/resolve", ctrl.ResolveTask)
	})

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

	task := r.Context().Value(taskContextkey).(*models.Task)

	dto := &api.UpdateTaskRequestDto{}
	if err := render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	// TODO validate user owns focus area
	update := models.Task{
		Model:       gorm.Model{ID: task.ID},
		Summary:     dto.Summary,
		Notes:       dto.Notes,
		FocusAreaID: dto.FocusAreaID,
	}

	updatedTask, err := ctrl.taskService.UpdateUserTask(ctx, &update)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Render(w, r, updatedTask.AsDTO())
}

func (ctrl *TaskController) ResolveTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	task := r.Context().Value(taskContextkey).(*models.Task)

	updatedTask, err := ctrl.taskService.ResolveUserTask(ctx, task.UserID, task.ID)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Render(w, r, updatedTask.AsDTO())
}

func (ctrl *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	task := r.Context().Value(taskContextkey).(*models.Task)

	err := ctrl.taskService.DeleteUserTask(ctx, task.ID)
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

func (ctrl *TaskController) taskContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		taskID := chi.URLParam(r, "taskID")
		if taskID == "" {
			render.Render(w, r, common.ErrNotFound)
			return
		}

		taskIDInt, err := strconv.Atoi(taskID)
		if err != nil {
			render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("failed to parse taskID: %w", err)))
			return
		}

		user, err := auth.GetUserFromCtx(ctx)
		if err != nil {
			render.Render(w, r, common.ErrUnauthorized(err))
			return
		}

		task, err := ctrl.taskService.GetUserTask(ctx, user.ID, uint(taskIDInt))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				render.Render(w, r, common.ErrNotFound)
			} else {
				render.Render(w, r, common.ErrUnknown(err))
			}

			return
		}

		ctxWithTask := context.WithValue(r.Context(), taskContextkey, &task)

		next.ServeHTTP(w, r.WithContext(ctxWithTask))
	})
}
