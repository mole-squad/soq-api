package api

import "net/http"

type TokenResponseDTO struct {
	Token string `json:"token"`
}

func NewTokenResponseDTO(token string) *TokenResponseDTO {
	dto := &TokenResponseDTO{
		Token: token,
	}

	return dto
}

func (t *TokenResponseDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
