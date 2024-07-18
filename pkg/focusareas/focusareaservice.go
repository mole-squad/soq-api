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
	*generics.ResourceService[*models.FocusArea]

	focusAreaRepo interfaces.FocusAreaRepo
}

func NewFocusAreaService(params FocusAreaServiceParams) (FocusAreaServiceResult, error) {
	embeddedSvc := generics.NewResourceService[*models.FocusArea](
		params.FocusAreaRepo,
	).(*generics.ResourceService[*models.FocusArea])

	srv := &FocusAreaService{
		ResourceService: embeddedSvc,
		focusAreaRepo:   params.FocusAreaRepo,
	}

	return FocusAreaServiceResult{FocusAreaService: srv}, nil
}
