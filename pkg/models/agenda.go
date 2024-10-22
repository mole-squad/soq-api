package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type AgendaStatus int

const (
	AgendaStatusPending AgendaStatus = iota
	AgendaStatusGenerated
	AgendaStatusSent
)

type Agenda struct {
	gorm.Model

	Status AgendaStatus

	StartTime time.Time
	EndTime   time.Time

	AgendaItems []AgendaItem

	FocusAreaID uint
	FocusArea   FocusArea `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (a *Agenda) GetID() uint {
	return a.ID
}

func (a *Agenda) GetTitle() string {
	return fmt.Sprintf(
		"%s Agenda %s",
		a.FocusArea.Name,
		a.StartTime.Format("Jan 1, 2006"),
	)
}

func (a *Agenda) GetBody() string {
	var builder strings.Builder

	for _, item := range a.AgendaItems {
		builder.WriteString(" - ")
		builder.WriteString(item.GetShortDescription())
		builder.WriteString("\n")
	}

	return builder.String()
}

func (a *Agenda) ToDTO() render.Renderer {
	return nil
}
