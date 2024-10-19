package interfaces

import (
	"github.com/mole-squad/soq-api/pkg/models"
)

type QuotaRepo interface {
	ResourceRepository[*models.Quota]
}
