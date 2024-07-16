package api

import (
	"fmt"
	"net/http"
)

type CreateDeviceRequestDTO struct {
	UserKey  string `json:"userKey"`
	DeviceID string `json:"deviceId"`
}

func (dto *CreateDeviceRequestDTO) Bind(r *http.Request) error {
	if dto.UserKey == "" {
		return fmt.Errorf("userKey is required")
	}

	if dto.DeviceID == "" {
		return fmt.Errorf("deviceID is required")
	}

	return nil
}
