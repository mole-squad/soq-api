package quotas

import "net/http"

type CreateQuotaRequestDTO struct {
	Summary     string `json:"summary"`
	Notes       string `json:"notes"`
	FocusAreaID uint   `json:"focusAreaId"`
}

func (dto *CreateQuotaRequestDTO) Bind(r *http.Request) error {
	return nil
}
