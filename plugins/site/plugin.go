package site

import (
	"github.com/kapmahc/h2o/lib/app"
	"github.com/kapmahc/h2o/lib/job"
	"github.com/kapmahc/h2o/lib/mux"
	"github.com/urfave/cli"
)

// Plugin plugin
type Plugin struct {
}

// Open open
func (p *Plugin) Open(*app.Config) error {
	return nil
}

// Routes http routes
func (p *Plugin) Routes() []mux.Route {
	return []mux.Route{}
}

// Workers background workers
func (p *Plugin) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

// Commands console commands
func (p *Plugin) Commands() []cli.Command {
	return []cli.Command{}
}

func init() {
	app.Register(&Plugin{})
}
