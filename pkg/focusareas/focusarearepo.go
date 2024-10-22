package focusareas

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type FocusAreaRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService mochi.LoggerService
}

type FocusAreaRepoResult struct {
	fx.Out

	FocusAreaRepo interfaces.FocusAreaRepo
}

type FocusAreaRepo struct {
	mochi.Repository[*models.FocusArea]
}

func NewFocusAreaRepo(params FocusAreaRepoParams) (FocusAreaRepoResult, error) {
	embeddedRepo := mochi.NewRepository(params.DBService, params.LoggerService, mochi.WithTableName[*models.FocusArea]("focus_areas"), mochi.WithPreloadTables[*models.FocusArea]("TimeWindows"))

	repo := &FocusAreaRepo{
		Repository: embeddedRepo,
	}

	return FocusAreaRepoResult{FocusAreaRepo: repo}, nil
}
