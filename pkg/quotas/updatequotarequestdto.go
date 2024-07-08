package quotas

import "net/http"

type UpdateQuotaRequestDto struct {
	// TODO
	Summary     string `json:"summary"`
	FocusAreaID uint   `json:"focusAreaId"`
}

func (dto *UpdateQuotaRequestDto) Bind(r *http.Request) error {
	return nil
}
