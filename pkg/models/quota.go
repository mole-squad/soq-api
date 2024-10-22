package models

import (
	"net/http"

	"github.com/burkel24/go-mochi"
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/api"
	"gorm.io/gorm"
)

type QuotaPeriod = int

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

func (q *Quota) GetID() uint {
	return q.ID
}

func (q *Quota) ToDTO() render.Renderer {
	focusArea := q.FocusArea.ToDTO().(*api.FocusAreaDTO)

	dto := &api.QuotaDTO{
		ID:        q.ID,
		Summary:   q.Summary,
		FocusArea: *focusArea,
	}

	return dto
}

func NewQuotaFromCreateRequest(r *http.Request, user mochi.User) (*Quota, error) {
	quota := &Quota{}

	dto := &api.CreateQuotaRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		return nil, err
	}

	quota.Summary = dto.Summary
	quota.TargetTimeMins = dto.TargetTimeMins
	quota.TargetInstances = dto.TargetInstances
	quota.Period = QuotaPeriod(dto.Period)

	return quota, nil
}

func NewQuotaFromUpdateRequest(r *http.Request, user mochi.User) (*Quota, error) {
	quota := &Quota{}

	dto := &api.UpdateQuotaRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		return nil, err
	}

	quota.Summary = dto.Summary
	quota.TargetTimeMins = dto.TargetTimeMins
	quota.TargetInstances = dto.TargetInstances
	quota.Period = QuotaPeriod(dto.Period)

	return quota, nil
}
