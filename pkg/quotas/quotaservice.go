package quotas

import (
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type QuotaServiceParams struct {
	fx.In

	QuotaRepo interfaces.QuotaRepo
}

type QuotaServiceResult struct {
	fx.Out

	QuotaService interfaces.QuotaService
}

type QuotaService struct {
	*generics.Service[*models.Quota]
}

func NewQuotaService(params QuotaServiceParams) (QuotaServiceResult, error) {
	embeddedSvc := generics.NewService[*models.Quota](
		params.QuotaRepo,
	).(*generics.Service[*models.Quota])

	srv := &QuotaService{
		Service: embeddedSvc,
	}

	return QuotaServiceResult{QuotaService: srv}, nil
}
