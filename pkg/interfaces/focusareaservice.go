package interfaces

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/models"
)

type FocusAreaService interface {
	mochi.Service[*models.FocusArea]
}
