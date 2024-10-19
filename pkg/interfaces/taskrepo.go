package interfaces

import "github.com/mole-squad/soq-api/pkg/models"

type TaskRepo interface {
	ResourceRepository[*models.Task]
}
