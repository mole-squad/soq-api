package quotas

import (
	"net/http"

	"github.com/mole-squad/soq-api/pkg/models"
)

type CreateQuotaRequestDTO struct {
	Summary string `json:"summary"`

	TargetTimeMins  int `json:"targetTimeMins"`
	TargetInstances int `json:"targetInstances"`

	Period models.QuotaPeriod `json:"period"`

	FocusAreaID uint `json:"focusAreaId"`
}

func (dto *CreateQuotaRequestDTO) Bind(r *http.Request) error {
	return nil
}
