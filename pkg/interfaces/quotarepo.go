package interfaces

import (
	"github.com/mole-squad/soq-api/pkg/models"
)

type QuotaRepo interface {
	Repository[*models.Quota]
}
