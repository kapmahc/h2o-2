package auth

import (
	"github.com/kapmahc/h2o/lib/app"
	"github.com/kapmahc/h2o/lib/mux"
)

// Plugin plugin
type Plugin struct {
}

// Routes http routes
func (p *Plugin) Routes() []mux.Route {
	return []mux.Route{}
}

func init() {
	app.Register(&Plugin{})
}
