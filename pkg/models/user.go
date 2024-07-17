package models

import (
	"github.com/mole-squad/soq-api/api"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	PasswordHash string

	Admin bool

	Name     string
	Timezone string

	Agendas     []Agenda
	FocusAreas  []FocusArea
	Quotas      []Quota
	Tasks       []Task
	TimeWindows []TimeWindow
}

func (u *User) ToDTO() *api.UserDTO {
	return &api.UserDTO{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		Timezone: u.Timezone,
	}
}
