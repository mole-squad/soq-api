package models

import (
	"github.com/lib/pq"
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
