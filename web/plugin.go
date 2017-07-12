package web

import (
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin interface {
	Mount(*gin.Engine)
	Shell() []cli.Command
	RegisterWorker()
	Atom(lang string) ([]*atom.Entry, error)
	Sitemap() ([]stm.URL, error)
	Dashboard(*gin.Context) *Dropdown
}

// -----------------------------------------------------------------------------

var plugins []Plugin

// Register register engines
func Register(ens ...Plugin) {
	plugins = append(plugins, ens...)
}

// Walk walk engines
func Walk(fn func(Plugin) error) error {
	for _, en := range plugins {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
