package models

import (
	"net/http"

	"github.com/go-chi/render"
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

func (d *Device) GetID() uint {
	return d.ID
}

func (d *Device) GetUserID() uint {
	return d.UserID
}

func (d *Device) SetUserID(userID uint) {
	d.UserID = userID
}

func (d *Device) ToDTO() render.Renderer {
	return &api.DeviceDTO{
		ID:       d.ID,
		UserKey:  d.UserKey,
		DeviceID: d.DeviceID,
	}
}

func NewDeviceFromCreateRequest(r *http.Request) (*Device, error) {
	device := &Device{}

	dto := &api.CreateDeviceRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		return nil, err
	}

	device.UserKey = dto.UserKey
	device.DeviceID = dto.DeviceID

	return device, nil
}

func NewDeviceFromUpdateRequest(r *http.Request) (*Device, error) {
	device := &Device{}

	dto := &api.UpdateDeviceRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		return nil, err
	}

	device.UserKey = dto.UserKey
	device.DeviceID = dto.DeviceID

	return device, nil
}
