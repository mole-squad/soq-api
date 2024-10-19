package api

import (
	"github.com/burkel24/go-mochi"
)

const (
	taskContextkey mochi.ResourceContextKey = iota
	deviceContextKey
	focusAreaContextKey
	quotaContextKey
	timeWindowContextKey
)
