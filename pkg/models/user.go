package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	PasswordHash string

	Name     string
	Timezone string

	Agendas     []Agenda
	FocusAreas  []FocusArea
	Quotas      []Quota
	Tasks       []Task
	TimeWindows []TimeWindow
}
