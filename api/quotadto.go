package api

import (
	"net/http"
)

type QuotaDTO struct {
	ID      uint   `json:"id"`
	Summary string `json:"summary"`

	TargetTimeMins  int `json:"targetTimeMins"`
	TargetInstances int `json:"targetInstances"`

	Period int `json:"period"`

	FocusArea FocusAreaDTO `json:"focusArea"`
}

func (t *QuotaDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
