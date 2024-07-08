package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Timezone string

	FocusAreas  []FocusArea
	Quotas      []Quota
	Tasks       []Task
	TimeWindows []TimeWindow
}
