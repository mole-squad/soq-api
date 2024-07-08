package models

import "gorm.io/gorm"

type QuotaPeriod int

const (
	DailyQuota   QuotaPeriod = iota
	WeeklyQuota              = iota
	MonthlyQuota             = iota
)

type Quota struct {
	gorm.Model

	Summary string

	TargetTimeMins  int
	TargetInstances int

	Period QuotaPeriod

	FocusAreaID uint
	FocusArea   FocusArea `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}