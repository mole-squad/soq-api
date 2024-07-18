package api

import "net/http"

type UpdateTimeWindowRequestDTO struct {
	Weekdays  []int32 `json:"weekdays"`
	StartTime float32 `json:"startTime"`
	EndTime   float32 `json:"endTime"`
}

func (dto *UpdateTimeWindowRequestDTO) Bind(r *http.Request) error {
	return nil
}
