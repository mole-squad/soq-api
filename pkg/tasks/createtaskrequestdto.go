package tasks

import "net/http"

type CreateTaskRequestDto struct {
	Summary string
}

func (dto *CreateTaskRequestDto) Bind(r *http.Request) error {
	return nil
}
