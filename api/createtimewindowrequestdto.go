package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateTimeWindowRequestDTO struct {
	Weekdays       []int32     `json:"weekdays"`
	StartTime      float32     `json:"startTime"`
	EndTime        float32     `json:"endTime"`
	FocusAreaIDNum json.Number `json:"focusAreaId"`
	FocusAreaID    uint
}

func (dto *CreateTimeWindowRequestDTO) Bind(r *http.Request) error {
	if dto.FocusAreaIDNum.String() == "" {
		return nil
	}

	focusAreaID, err := dto.FocusAreaIDNum.Int64()
	if err != nil {
		return fmt.Errorf("failed to parse focusAreaId: %w", err)
	}

	dto.FocusAreaID = uint(focusAreaID)

	return nil
}
