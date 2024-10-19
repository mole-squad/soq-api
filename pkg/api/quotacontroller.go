package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type QuotaControllerParams struct {
	fx.In

	AuthService  interfaces.AuthService
	Logger       interfaces.LoggerService
	QuotaService interfaces.QuotaService
	Router       *chi.Mux
}

type QuotaControllerResult struct {
	fx.Out

	QuotaController QuotaController
}

type QuotaController struct {
	interfaces.ResourceController[*models.Quota]
}

func NewQuotaController(params QuotaControllerParams) (QuotaControllerResult, error) {
	ctrl := QuotaController{}

	ctrl.ResourceController = generics.NewController[*models.Quota](
		params.QuotaService,
		params.Logger,
		params.AuthService,
		models.NewQuotaFromCreateRequest,
		models.NewQuotaFromUpdateRequest,
		generics.WithContextKey[*models.Quota](quotaContextKey),
	).(*generics.Controller[*models.Quota])

	params.Router.Mount("/quotas", ctrl.ResourceController.GetRouter())

	return QuotaControllerResult{QuotaController: ctrl}, nil
}
