package interfaces

import (
	"github.com/mole-squad/soq-api/pkg/models"
)

type QuotaService interface {
	Service[*models.Quota]
}
