package interfaces

import (
	"github.com/mole-squad/soq-api/pkg/models"
)

type FocusAreaRepo interface {
	ResourceRepository[*models.FocusArea]
}
