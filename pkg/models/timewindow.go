package models

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/lib/pq"
	"github.com/mole-squad/soq-api/api"
	"gorm.io/gorm"
)

type TimeWindow struct {
	gorm.Model

	Weekdays pq.Int32Array `gorm:"type:int[];"`

	StartTime float32
	EndTime   float32

	FocusAreaID uint
	FocusArea   FocusArea `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (t *TimeWindow) GetID() uint {
	return t.ID
}

func (t *TimeWindow) GetUserID() uint {
	return t.UserID
}

func (t *TimeWindow) SetUserID(userID uint) {
	t.UserID = userID
}

func (t *TimeWindow) ToDTO() render.Renderer {
	dto := &api.TimeWindowDTO{
		ID:        t.ID,
		Weekdays:  t.Weekdays,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
	}

	return dto
}

func NewTimeWindowFromCreateRequest(r *http.Request) (*TimeWindow, error) {
	timeWindow := &TimeWindow{}

	dto := &api.CreateTimeWindowRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		return nil, err
	}

	timeWindow.Weekdays = dto.Weekdays
	timeWindow.StartTime = dto.StartTime
	timeWindow.EndTime = dto.EndTime
	timeWindow.FocusAreaID = dto.FocusAreaID

	return timeWindow, nil
}

func NewTimeWindowFromUpdateRequest(r *http.Request) (*TimeWindow, error) {
	timeWindow := &TimeWindow{}

	dto := &api.UpdateTimeWindowRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		return nil, err
	}

	timeWindow.Weekdays = dto.Weekdays
	timeWindow.StartTime = dto.StartTime
	timeWindow.EndTime = dto.EndTime

	return timeWindow, nil
}
