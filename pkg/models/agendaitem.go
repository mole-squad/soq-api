package models

import "gorm.io/gorm"

type AgendaItem struct {
	gorm.Model

	AgendaID uint
	Agenda   Agenda `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `

	TaskId *uint `gorm:"default:null"`
	Task   *Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	QuotaID *uint  `gorm:"default:null"`
	Quota   *Quota `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
