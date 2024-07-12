package users

import "net/http"

type SetPasswordRequestDTO struct {
	Password string `json:"password"`
}

func (dto *SetPasswordRequestDTO) Bind(r *http.Request) error {
	return nil
}
