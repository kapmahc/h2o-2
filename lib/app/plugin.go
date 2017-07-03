package app

import (
	"github.com/kapmahc/air/web/job"
	"github.com/kapmahc/h2o/lib/mux"
)

// Plugin plugin
type Plugin interface {
	Open(*Config) error
	Routes() []mux.Route
	Workers() map[string]job.Handler
}

var plugins []Plugin

// Register register plugins
func Register(items ...Plugin) {
	plugins = append(plugins, items...)
}

// Loop loop plugins
func Loop(f func(Plugin) error) error {
	for _, p := range plugins {
		if e := f(p); e != nil {
			return e
		}
	}
	return nil
}
