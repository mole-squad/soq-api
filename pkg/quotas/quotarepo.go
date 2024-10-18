package quotas

import (
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type QuotaRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService interfaces.LoggerService
}

type QuotaRepoResult struct {
	fx.Out

	QuotaRepo interfaces.QuotaRepo
}

type QuotaRepo struct {
	*generics.Repository[*models.Quota]
}

func NewQuotaRepo(params QuotaRepoParams) (QuotaRepoResult, error) {
	embeddedRepo := generics.NewRepository[*models.Quota](
		params.DBService,
		params.LoggerService,
		generics.WithTableName[*models.Quota]("quotas"),
		generics.WithJoinTables[*models.Quota]("FocusArea"),
	).(*generics.Repository[*models.Quota])

	repo := &QuotaRepo{
		Repository: embeddedRepo,
	}

	return QuotaRepoResult{QuotaRepo: repo}, nil
}
