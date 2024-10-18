package focusareas

import (
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type FocusAreaRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService interfaces.LoggerService
}

type FocusAreaRepoResult struct {
	fx.Out

	FocusAreaRepo interfaces.FocusAreaRepo
}

type FocusAreaRepo struct {
	*generics.Repository[*models.FocusArea]
}

func NewFocusAreaRepo(params FocusAreaRepoParams) (FocusAreaRepoResult, error) {
	embeddedRepo := generics.NewRepository[*models.FocusArea](
		params.DBService,
		params.LoggerService,
		generics.WithTableName[*models.FocusArea]("focus_areas"),
		generics.WithPreloadTables[*models.FocusArea]("TimeWindows"),
	).(*generics.Repository[*models.FocusArea])

	repo := &FocusAreaRepo{
		Repository: embeddedRepo,
	}

	return FocusAreaRepoResult{FocusAreaRepo: repo}, nil
}
