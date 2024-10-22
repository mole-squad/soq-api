package interfaces

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/models"
)

type DeviceRepo interface {
	mochi.Repository[*models.Device]
}
