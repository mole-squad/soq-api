package quotas

import (
	"github.com/burkel24/go-mochi"
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
	mochi.Service[*models.Quota]
}

func NewQuotaService(params QuotaServiceParams) (QuotaServiceResult, error) {
	embeddedSvc := mochi.NewService(
		params.QuotaRepo,
	)

	srv := &QuotaService{
		Service: embeddedSvc,
	}

	return QuotaServiceResult{QuotaService: srv}, nil
}
