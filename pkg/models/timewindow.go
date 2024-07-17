package models

import (
	"github.com/lib/pq"
	"github.com/mole-squad/soq-api/api"
	"gorm.io/gorm"
)

type TimeWindow struct {
	gorm.Model

	Weekdays pq.Int32Array `gorm:"type:int[];"`

	StartTime float32
	EndTime   float32

	FocusAreaID uint
	FocusArea   FocusArea `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (t *TimeWindow) ToDTO() *api.TimeWindowDTO {
	dto := &api.TimeWindowDTO{
		ID:        t.ID,
		Weekdays:  t.Weekdays,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
	}

	return dto
}
