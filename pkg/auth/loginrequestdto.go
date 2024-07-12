package auth

import "net/http"

type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto *LoginRequestDTO) Bind(r *http.Request) error {
	return nil
}
