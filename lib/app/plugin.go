package app

import (
	"github.com/kapmahc/h2o/lib/job"
	"github.com/kapmahc/h2o/lib/mux"
	"github.com/urfave/cli"
)

// Plugin plugin
type Plugin interface {
	Open(*Config) error
	Routes() []mux.Route
	Workers() map[string]job.Handler
	Commands() []cli.Command
}

var plugins []Plugin

// Register register plugins
func Register(items ...Plugin) {
	plugins = append(plugins, items...)
}

// Walk walk plugins
func Walk(f func(Plugin) error) error {
	for _, p := range plugins {
		if e := f(p); e != nil {
			return e
		}
	}
	return nil
}
