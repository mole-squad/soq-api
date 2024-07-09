package models

import (
	"time"

	"gorm.io/gorm"
)

type AgendaStatus int

const (
	AgendaStatusPending AgendaStatus = iota
	AgendaStatusCompleted
)

type Agenda struct {
	gorm.Model

	Status AgendaStatus

	StartTime time.Time
	EndTime   time.Time

	AgendaItems []AgendaItem

	FocusAreaID uint
	FocusArea   FocusArea `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
