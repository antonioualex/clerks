package domain

import "net/http"

type RouteDef struct {
	Methods     []string
	Handler     http.Handler
	HandlerFunc http.HandlerFunc
}
