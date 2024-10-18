package interfaces

import (
	"github.com/mole-squad/soq-api/pkg/models"
)

type DeviceRepo interface {
	Repository[*models.Device]
}
