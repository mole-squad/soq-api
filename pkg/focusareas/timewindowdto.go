package focusareas

import (
	"net/http"

	"github.com/burkel24/task-app/pkg/models"
)

type TimeWindowDTO struct {
	ID uint `json:"id"`

	Weekdays []int32 `json:"weekdays"`

	StartTime float32 `json:"startTime"`
	EndTime   float32 `json:"endTime"`
}

func NewTimeWindowDTO(timeWindow models.TimeWindow) *TimeWindowDTO {
	dto := &TimeWindowDTO{
		ID:        timeWindow.ID,
		Weekdays:  timeWindow.Weekdays,
		StartTime: timeWindow.StartTime,
		EndTime:   timeWindow.EndTime,
	}

	return dto
}

func (t *TimeWindowDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
