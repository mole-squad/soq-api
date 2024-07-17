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

type QuotaControllerParams struct {
	fx.In

	AuthService  interfaces.AuthService
	QuotaService interfaces.QuotaService
	Router       *chi.Mux
}

type QuotaControllerResult struct {
	fx.Out

	QuotaController QuotaController
}

type QuotaController struct {
	quotaService interfaces.QuotaService
}

func NewQuotaController(params QuotaControllerParams) (QuotaControllerResult, error) {
	ctrl := QuotaController{quotaService: params.QuotaService}

	quotaRouter := chi.NewRouter()
	quotaRouter.Use(params.AuthService.AuthRequired())

	quotaRouter.Get("/", ctrl.ListQuotas)
	quotaRouter.Post("/", ctrl.CreateQuota)
	quotaRouter.Patch("/{quotaID}", ctrl.UpdateQuota)
	quotaRouter.Delete("/{quotaID}", ctrl.DeleteQuota)

	params.Router.Mount("/quotas", quotaRouter)

	return QuotaControllerResult{QuotaController: ctrl}, nil
}

func (ctrl *QuotaController) CreateQuota(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	dto := &api.CreateQuotaRequestDTO{}
	if err = render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	// TODO validate user owns focus area
	newQuota := models.Quota{
		Summary:         dto.Summary,
		TargetTimeMins:  dto.TargetTimeMins,
		TargetInstances: dto.TargetInstances,
		Period:          dto.Period,
		FocusAreaID:     dto.FocusAreaID,
	}

	quota, err := ctrl.quotaService.CreateUserQuota(ctx, user, &newQuota)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Render(w, r, quota.ToDTO())
}

func (ctrl *QuotaController) UpdateQuota(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	quotaId := chi.URLParam(r, "quotaID")
	quotaIdInt, err := strconv.Atoi(quotaId)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("failed to parse quotaID: %w", err)))
		return
	}

	dto := &api.UpdateQuotaRequestDto{}
	if err = render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	// TODO validate user owns focus area
	quota := models.Quota{
		Model:           gorm.Model{ID: uint(quotaIdInt)},
		Summary:         dto.Summary,
		TargetTimeMins:  dto.TargetTimeMins,
		TargetInstances: dto.TargetInstances,
		Period:          dto.Period,
		FocusAreaID:     dto.FocusAreaID,
	}

	updatedQuota, err := ctrl.quotaService.UpdateUserQuota(ctx, &quota)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Render(w, r, updatedQuota.ToDTO())
}

func (ctrl *QuotaController) DeleteQuota(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	quotaId := chi.URLParam(r, "quotaID")
	quotaIdInt, err := strconv.Atoi(quotaId)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("failed to parse quotaID: %w", err)))
		return
	}

	err = ctrl.quotaService.DeleteUserQuota(ctx, uint(quotaIdInt))
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.NoContent(w, r)
}

func (ctrl *QuotaController) ListQuotas(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	userQuotas, err := ctrl.quotaService.ListUserQuotas(ctx, user)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	respList := []render.Renderer{}
	for _, quota := range userQuotas {
		respList = append(respList, quota.ToDTO())
	}

	render.RenderList(w, r, respList)
}
