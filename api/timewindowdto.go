package api

import (
	"net/http"
)

type TimeWindowDTO struct {
	ID uint `json:"id"`

	Weekdays []int32 `json:"weekdays"`

	StartTime float32 `json:"startTime"`
	EndTime   float32 `json:"endTime"`
}

func (t *TimeWindowDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
