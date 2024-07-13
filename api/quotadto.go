package api

import (
	"net/http"

	"github.com/mole-squad/soq-api/pkg/models"
)

type QuotaDTO struct {
	ID      uint   `json:"id"`
	Summary string `json:"summary"`

	TargetTimeMins  int `json:"targetTimeMins"`
	TargetInstances int `json:"targetInstances"`

	Period models.QuotaPeriod `json:"period"`

	FocusArea FocusAreaDTO `json:"focusArea"`
}

func NewQuotaDTO(quota models.Quota) *QuotaDTO {
	// TODO
	dto := &QuotaDTO{
		ID:        quota.ID,
		Summary:   quota.Summary,
		FocusArea: *NewFocusAreaDTO(quota.FocusArea),
	}

	return dto
}

func (t *QuotaDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
