package site

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/h2o/plugins/auth"
	"github.com/kapmahc/h2o/web"
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Server   *machinery.Server `inject:""`
	Cache    *web.Cache        `inject:""`
	I18n     *web.I18n         `inject:""`
	Settings *web.Settings     `inject:""`
	Db       *gorm.DB          `inject:""`
	Jwt      *auth.Jwt         `inject:""`
	Redis    *redis.Pool       `inject:""`
	Matcher  language.Matcher  `inject:""`
	Dao      *auth.Dao         `inject:""`
}

// RegisterWorker register worker
func (p *Plugin) RegisterWorker() {

}

// Atom rss.atom
func (p *Plugin) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Plugin) Sitemap() ([]stm.URL, error) {
	urls := []stm.URL{
		{"loc": "/"},
		{"loc": "/leave-words/new"},
		{"loc": "/notices"},
	}
	return urls, nil
}

// Dashboard dashboard
func (p *Plugin) Dashboard(c *gin.Context) *web.Dropdown {
	if admin, ok := c.Get(auth.IsAdmin); !ok || !admin.(bool) {
		return nil
	}
	return &web.Dropdown{
		Label: "site.dashboard.title",
		Links: []*web.Link{
			&web.Link{Href: "/admin/site/status", Label: "site.admin.status.title"},
			nil,
			&web.Link{Href: "/admin/links", Label: "site.admin.links.index.title"},
			&web.Link{Href: "/admin/pages", Label: "site.admin.pages.index.title"},
			nil,
			&web.Link{Href: "/admin/site/info", Label: "site.admin.info.title"},
			&web.Link{Href: "/admin/site/author", Label: "site.admin.author.title"},
			&web.Link{Href: "/admin/site/seo", Label: "site.admin.seo.title"},
			&web.Link{Href: "/admin/site/smtp", Label: "site.admin.smtp.title"},
			nil,
			&web.Link{Href: "/admin/users", Label: "site.admin.users.index.title"},
			&web.Link{Href: "/admin/locales", Label: "site.admin.locales.index.title"},
			&web.Link{Href: "/admin/notices", Label: "site.notices.index.title"},
			&web.Link{Href: "/leave-words", Label: "site.leave-words.index.title"},
		},
	}
}

func init() {
	web.Register(&Plugin{})
}
