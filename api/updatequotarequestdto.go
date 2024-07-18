package api

import (
	"net/http"
)

type UpdateQuotaRequestDTO struct {
	Summary string `json:"summary"`

	TargetTimeMins  int `json:"targetTimeMins"`
	TargetInstances int `json:"targetInstances"`

	Period int `json:"period"`

	FocusAreaID uint `json:"focusAreaId"`
}

func (dto *UpdateQuotaRequestDTO) Bind(r *http.Request) error {
	// TODO validate user owns focusarea

	return nil
}
