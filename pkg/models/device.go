package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model

	UserKey  string
	DeviceID string

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
