package mux

// Handler http handler
type Handler func(*Context) error

// Route http route
type Route struct {
	Method  string
	Path    string
	Handler Handler
}
