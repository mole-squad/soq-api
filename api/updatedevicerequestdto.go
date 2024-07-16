package api

import (
	"fmt"
	"net/http"
)

type UpdateDeviceRequestDTO struct {
	UserKey  string `json:"userKey"`
	DeviceID string `json:"deviceId"`
}

func (dto *UpdateDeviceRequestDTO) Bind(r *http.Request) error {
	if dto.UserKey == "" {
		return fmt.Errorf("userKey is required")
	}

	if dto.DeviceID == "" {
		return fmt.Errorf("deviceID is required")
	}

	return nil
}
