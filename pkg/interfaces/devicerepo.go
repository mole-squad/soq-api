package interfaces

import (
	"github.com/mole-squad/soq-api/pkg/models"
)

type DeviceRepo interface {
	ResourceRepository[*models.Device]
}
