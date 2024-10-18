package interfaces

import "github.com/mole-squad/soq-api/pkg/models"

type TimeWindowRepo interface {
	Repository[*models.TimeWindow]
}
