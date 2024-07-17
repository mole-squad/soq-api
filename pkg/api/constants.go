package api

import "github.com/mole-squad/soq-api/pkg/rest"

type contextKey int

const (
	taskContextkey rest.ResourceContextKey = iota
	deviceContextKey
)
