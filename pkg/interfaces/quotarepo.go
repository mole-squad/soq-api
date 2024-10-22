package interfaces

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/models"
)

type QuotaRepo interface {
	mochi.Repository[*models.Quota]
}
