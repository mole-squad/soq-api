package interfaces

import (
	"github.com/mole-squad/soq-api/pkg/models"
)

type TimeWindowService interface {
	Service[*models.TimeWindow]
}
