package api

import (
	"net/http"
)

type FocusAreaDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	TimeWindows []TimeWindowDTO `json:"timeWindows"`
}

func (f *FocusAreaDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
