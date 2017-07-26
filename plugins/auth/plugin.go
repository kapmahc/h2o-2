package auth

import (
	"html/template"
	"path"

	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/base"
	"github.com/kapmahc/axe/cache"
	"github.com/kapmahc/axe/i18n"
	"github.com/kapmahc/axe/job"
	"github.com/kapmahc/axe/settings"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Db       *gorm.DB           `inject:""`
	I18n     *i18n.I18n         `inject:""`
	Settings *settings.Settings `inject:""`
	Server   *job.Server        `inject:""`
	Cache    *cache.Cache       `inject:""`
	Wrapper  *axe.Wrapper       `inject:""`
}

// Workers workers
func (p *Plugin) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

// Rss rss.atom
func (p *Plugin) Rss(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Plugin) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Console console commands
func (p *Plugin) Console() []cli.Command {
	return []cli.Command{}
}

// Open open beans
func (p *Plugin) Open(g *inject.Graph) error {
	theme := viper.GetString("server.theme")
	prod := base.IsProduction()
	rdr := render.New(render.Options{
		Directory:  path.Join("themes", theme, "views"),
		Extensions: []string{".html"},
		Funcs:      []template.FuncMap{},
		IndentJSON: !prod,
		IndentXML:  !prod,
	})
	return g.Provide(
		&inject.Object{Value: rdr},
	)
}

func init() {
	axe.Register(&Plugin{})
}
