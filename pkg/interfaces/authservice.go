package interfaces

import "net/http"

type AuthService interface {
	AuthRequired() func(http.Handler) http.Handler
}
