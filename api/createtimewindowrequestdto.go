package api

import "net/http"

type CreateTimeWindowRequestDTO struct {
	Weekdays    []int32 `json:"weekdays"`
	StartTime   float32 `json:"startTime"`
	EndTime     float32 `json:"endTime"`
	FocusAreaID uint    `json:"focusAreaId"`
}

func (dto *CreateTimeWindowRequestDTO) Bind(r *http.Request) error {
	return nil
}
