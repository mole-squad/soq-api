package models

import "gorm.io/gorm"

type AgendaItem struct {
	gorm.Model

	AgendaID uint
	Agenda   Agenda `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `

	TaskID *uint `gorm:"default:null"`
	Task   *Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	QuotaID *uint  `gorm:"default:null"`
	Quota   *Quota `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (ai *AgendaItem) GetShortDescription() string {
	if ai.TaskID != nil {
		return ai.getTaskShortDescription()
	}

	if ai.QuotaID != nil {
		return ai.getQuotaShortDescription()
	}

	return "Unknown Agenda Item"
}

func (ai *AgendaItem) getTaskShortDescription() string {
	return ai.Task.Summary
}

func (ai *AgendaItem) getQuotaShortDescription() string {
	return "TODO"
}
