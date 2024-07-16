package api

import "net/http"

type DeviceDTO struct {
	ID uint `json:"id"`

	UserKey  string `json:"userKey"`
	DeviceID string `json:"deviceID"`
}

func (t *DeviceDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
