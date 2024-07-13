package models

import (
	"github.com/mole-squad/soq-api/api"
	"gorm.io/gorm"
)

type FocusArea struct {
	gorm.Model

	Name string

	TimeWindows []TimeWindow

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (f *FocusArea) AsDTO() *api.FocusAreaDTO {
	timeWindows := make([]api.TimeWindowDTO, len(f.TimeWindows))
	for i, timeWindow := range f.TimeWindows {
		timeWindows[i] = *timeWindow.AsDTO()
	}

	dto := &api.FocusAreaDTO{
		ID:          f.ID,
		Name:        f.Name,
		TimeWindows: timeWindows,
	}

	return dto
}
