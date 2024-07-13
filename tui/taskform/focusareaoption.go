package taskform

import "github.com/mole-squad/soq-api/pkg/focusareas"

type focusAreaOption struct {
	focusArea focusareas.FocusAreaDTO
}

func NewFocusAreaOption(fa focusareas.FocusAreaDTO) *focusAreaOption {
	return &focusAreaOption{
		focusArea: fa,
	}
}

func (f *focusAreaOption) Label() string {
	return f.focusArea.Name
}

func (f *focusAreaOption) Value() interface{} {
	return f.focusArea.ID
}
