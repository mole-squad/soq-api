package models

import "gorm.io/gorm"

type FocusArea struct {
	gorm.Model

	Name   string
	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
