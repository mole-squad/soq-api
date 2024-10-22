package quotas

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type QuotaRepoParams struct {
	fx.In

	DBService     mochi.DBService
	LoggerService mochi.LoggerService
}

type QuotaRepoResult struct {
	fx.Out

	QuotaRepo interfaces.QuotaRepo
}

type QuotaRepo struct {
	mochi.Repository[*models.Quota]
}

func NewQuotaRepo(params QuotaRepoParams) (QuotaRepoResult, error) {
	embeddedRepo := mochi.NewRepository(
		params.DBService,
		params.LoggerService,
		mochi.WithTableName[*models.Quota]("quotas"),
		mochi.WithJoinTables[*models.Quota]("FocusArea"),
	)

	repo := &QuotaRepo{
		Repository: embeddedRepo,
	}

	return QuotaRepoResult{QuotaRepo: repo}, nil
}
