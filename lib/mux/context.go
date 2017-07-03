package mux

import "net/http"

// Context http context
type Context struct {
	req *http.Request
	wrt http.ResponseWriter
}
