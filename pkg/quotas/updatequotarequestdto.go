package quotas

import (
	"net/http"

	"github.com/burkel24/task-app/pkg/models"
)

type UpdateQuotaRequestDto struct {
	Summary string `json:"summary"`

	TargetTimeMins  int `json:"targetTimeMins"`
	TargetInstances int `json:"targetInstances"`

	Period models.QuotaPeriod `json:"period"`

	FocusAreaID uint `json:"focusAreaId"`
}

func (dto *UpdateQuotaRequestDto) Bind(r *http.Request) error {
	return nil
}
