package interfaces

import (
	"github.com/mole-squad/soq-api/pkg/models"
)

type DeviceService interface {
	Service[*models.Device]
}
