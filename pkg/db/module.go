package db

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

func ProvideModels() mochi.ModelList {
	return []interface{}{
		&models.User{},
		&models.TimeWindow{},
		&models.FocusArea{},
		&models.Task{},
		&models.Quota{},
		&models.Agenda{},
		&models.AgendaItem{},
		&models.Device{},
	}
}

var Module = fx.Module(
	"DB",
	fx.Provide(mochi.NewDBService),
	fx.Provide(ProvideModels),
)
