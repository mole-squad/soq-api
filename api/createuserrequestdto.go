package api

import (
	"fmt"
	"net/http"
)

type CreateUserRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`

	Name     string `json:"name"`
	Timezone string `json:"timezone"`
}

func (dto *CreateUserRequestDTO) Bind(r *http.Request) error {
	if dto.Username == "" {
		return fmt.Errorf("username is required")
	}

	if dto.Password == "" {
		return fmt.Errorf("password is required")
	}

	if dto.Name == "" {
		return fmt.Errorf("name is required")
	}

	if dto.Timezone == "" {
		dto.Timezone = "America/Los_Angeles"
	}

	return nil
}
