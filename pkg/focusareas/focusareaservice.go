package focusareas

import (
	"github.com/mole-squad/soq-api/pkg/generics"
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
	*generics.Service[*models.FocusArea]
}

func NewFocusAreaService(params FocusAreaServiceParams) (FocusAreaServiceResult, error) {
	embeddedSvc := generics.NewService[*models.FocusArea](
		params.FocusAreaRepo,
	).(*generics.Service[*models.FocusArea])

	srv := &FocusAreaService{
		Service: embeddedSvc,
	}

	return FocusAreaServiceResult{FocusAreaService: srv}, nil
}
