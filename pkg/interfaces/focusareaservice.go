package interfaces

import (
	"github.com/mole-squad/soq-api/pkg/models"
)

type FocusAreaService interface {
	ResourceService[*models.FocusArea]
}
