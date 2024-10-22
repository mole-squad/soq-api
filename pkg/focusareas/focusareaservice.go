package focusareas

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type FocusAreaServiceParams struct {
	fx.In

	FocusAreaRepo interfaces.FocusAreaRepo
}

type FocusAreaServiceResult struct {
	fx.Out

	FocusAreaService interfaces.FocusAreaService
}

type FocusAreaService struct {
	mochi.Service[*models.FocusArea]
}

func NewFocusAreaService(params FocusAreaServiceParams) (FocusAreaServiceResult, error) {
	embeddedSvc := mochi.NewService(
		params.FocusAreaRepo,
	)

	srv := &FocusAreaService{
		Service: embeddedSvc,
	}

	return FocusAreaServiceResult{FocusAreaService: srv}, nil
}
