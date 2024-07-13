package api

import (
	"net/http"
)

type UpdateQuotaRequestDto struct {
	Summary string `json:"summary"`

	TargetTimeMins  int `json:"targetTimeMins"`
	TargetInstances int `json:"targetInstances"`

	Period int `json:"period"`

	FocusAreaID uint `json:"focusAreaId"`
}

func (dto *UpdateQuotaRequestDto) Bind(r *http.Request) error {
	return nil
}
