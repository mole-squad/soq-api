package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name string
}

func (u *User) Name() string {
	return u.name
}

func (u *User) ID() uint {
	return u.Model.ID
}
