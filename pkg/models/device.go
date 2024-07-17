package models

import (
	"github.com/mole-squad/soq-api/api"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model

	UserKey  string
	DeviceID string

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (d *Device) ToDTO() *api.DeviceDTO {
	return &api.DeviceDTO{
		ID:       d.ID,
		UserKey:  d.UserKey,
		DeviceID: d.DeviceID,
	}
}
