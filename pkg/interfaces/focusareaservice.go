package interfaces

import (
	"github.com/mole-squad/soq-api/pkg/models"
)

type FocusAreaService interface {
	Service[*models.FocusArea]
}
