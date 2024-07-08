package quotas

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/burkel24/task-app/pkg/auth"
	"github.com/burkel24/task-app/pkg/common"
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type QuotaControllerParams struct {
	fx.In

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

	dto := &CreateQuotaRequestDTO{}
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

	quota, err := ctrl.quotaService.CreateUserQuota(ctx, &user, &newQuota)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
	}

	resp := NewQuotaDTO(quota)
	render.Render(w, r, resp)
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
	}

	dto := &UpdateQuotaRequestDto{}
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
	}

	resp := NewQuotaDTO(updatedQuota)
	render.Render(w, r, resp)
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
	}

	err = ctrl.quotaService.DeleteUserQuota(ctx, uint(quotaIdInt))
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
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

	quotas, err := ctrl.quotaService.ListUserQuotas(ctx, &user)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
	}

	render.RenderList(w, r, NewQuotaListResponseDTO(quotas))
}