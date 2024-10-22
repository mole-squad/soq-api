package api

import (
	"github.com/burkel24/go-mochi"

	"github.com/go-chi/chi/v5"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type QuotaControllerParams struct {
	fx.In

	AuthService  mochi.AuthService
	Logger       mochi.LoggerService
	QuotaService interfaces.QuotaService
	Router       *chi.Mux
}

type QuotaControllerResult struct {
	fx.Out

	QuotaController QuotaController
}

type QuotaController struct {
	mochi.Controller[*models.Quota]
}

func NewQuotaController(params QuotaControllerParams) (QuotaControllerResult, error) {
	ctrl := QuotaController{}

	ctrl.Controller = mochi.NewController(
		params.QuotaService,
		params.Logger,
		params.AuthService,
		models.NewQuotaFromCreateRequest,
		models.NewQuotaFromUpdateRequest,
		mochi.WithContextKey[*models.Quota](quotaContextKey),
	)

	params.Router.Mount("/quotas", ctrl.Controller.GetRouter())

	return QuotaControllerResult{QuotaController: ctrl}, nil
}
