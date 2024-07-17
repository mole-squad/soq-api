package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/api"
	"github.com/mole-squad/soq-api/pkg/common"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type FocusAreaControllerParams struct {
	fx.In

	AuthService      interfaces.AuthService
	FocusAreaService interfaces.FocusAreaService
	Router           *chi.Mux
}

type FocusAreaControllerResult struct {
	fx.Out

	FocusAreaController FocusAreaController
}

type FocusAreaController struct {
	auth             interfaces.AuthService
	focusAreaService interfaces.FocusAreaService
}

func NewFocusAreaController(params FocusAreaControllerParams) (FocusAreaControllerResult, error) {
	ctrl := FocusAreaController{
		auth:             params.AuthService,
		focusAreaService: params.FocusAreaService,
	}

	focusAreaRouter := chi.NewRouter()

	focusAreaRouter.Use(params.AuthService.AuthRequired())

	focusAreaRouter.Get("/", ctrl.ListFocusAreas)
	focusAreaRouter.Post("/", ctrl.CreateFocusArea)
	focusAreaRouter.Patch("/{focusAreaID}", ctrl.UpdateFocusArea)
	focusAreaRouter.Delete("/{focusAreaID}", ctrl.DeleteFocusArea)

	params.Router.Mount("/focusareas", focusAreaRouter)

	return FocusAreaControllerResult{FocusAreaController: ctrl}, nil
}

func (ctrl *FocusAreaController) CreateFocusArea(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := ctrl.auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	dto := &api.CreateFocusAreaRequestDTO{}
	if err = render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	newFocusArea := &models.FocusArea{
		Name: dto.Name,
	}

	createdFocusArea, err := ctrl.focusAreaService.CreateFocusArea(ctx, user, newFocusArea)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Render(w, r, createdFocusArea.ToDTO())
}

func (ctrl *FocusAreaController) UpdateFocusArea(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := ctrl.auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	focusAreaID := chi.URLParam(r, "focusAreaID")
	focusAreaIDInt, err := strconv.Atoi(focusAreaID)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("failed to parse focusAreaID: %w", err)))
		return
	}

	dto := &api.UpdateFocusAreaRequestDTO{}
	if err = render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	focusArea := &models.FocusArea{
		Model: gorm.Model{ID: uint(focusAreaIDInt)},
		Name:  dto.Name,
	}

	updatedFocusArea, err := ctrl.focusAreaService.UpdateFocusArea(ctx, focusArea)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Render(w, r, updatedFocusArea.ToDTO())
}

func (ctrl *FocusAreaController) DeleteFocusArea(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := ctrl.auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	focusAreaID := chi.URLParam(r, "focusAreaID")
	focusAreaIDInt, err := strconv.Atoi(focusAreaID)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("failed to parse focusAreaID: %w", err)))
		return
	}

	err = ctrl.focusAreaService.DeleteFocusArea(ctx, uint(focusAreaIDInt))
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.NoContent(w, r)
}

func (ctrl *FocusAreaController) ListFocusAreas(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := ctrl.auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	focusAreas, err := ctrl.focusAreaService.ListUserFocusAreas(ctx, user)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	respList := []render.Renderer{}
	for _, focusArea := range focusAreas {
		respList = append(respList, focusArea.ToDTO())
	}

	render.RenderList(w, r, respList)
}
