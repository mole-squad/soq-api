package models

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/api"
	"gorm.io/gorm"
)

type FocusArea struct {
	gorm.Model

	Name string

	TimeWindows []TimeWindow

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (f *FocusArea) GetID() uint {
	return f.ID
}

func (f *FocusArea) GetUserID() uint {
	return f.UserID
}

func (f *FocusArea) SetUserID(userID uint) {
	f.UserID = userID
}

func (f *FocusArea) ToDTO() render.Renderer {
	timeWindows := make([]api.TimeWindowDTO, len(f.TimeWindows))
	for i, timeWindow := range f.TimeWindows {
		timeWindows[i] = *timeWindow.ToDTO().(*api.TimeWindowDTO)
	}

	dto := &api.FocusAreaDTO{
		ID:          f.ID,
		Name:        f.Name,
		TimeWindows: timeWindows,
	}

	return dto
}

func NewFocusAreaFromCreateRequest(r *http.Request) (*FocusArea, error) {
	focusArea := &FocusArea{}

	dto := &api.CreateFocusAreaRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		return nil, err
	}

	focusArea.Name = dto.Name

	return focusArea, nil
}

func NewFocusAreaFromUpdateRequest(r *http.Request) (*FocusArea, error) {
	focusArea := &FocusArea{}

	dto := &api.UpdateFocusAreaRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		return nil, err
	}

	focusArea.Name = dto.Name

	return focusArea, nil
}
