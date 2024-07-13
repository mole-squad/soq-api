package api

import "net/http"

type UserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
}

func (t *UserDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
