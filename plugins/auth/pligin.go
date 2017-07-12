package auth

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/h2o/web"
	"golang.org/x/tools/blog/atom"
)

// Plugin Plugin
type Plugin struct {
	Dao      *Dao              `inject:""`
	Db       *gorm.DB          `inject:""`
	Security *web.Security     `inject:""`
	I18n     *web.I18n         `inject:""`
	Jwt      *Jwt              `inject:""`
	Server   *machinery.Server `inject:""`
	Uploader web.Uploader      `inject:""`
	Settings *web.Settings     `inject:""`
}

// Atom rss.atom
func (p *Plugin) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Plugin) Sitemap() ([]stm.URL, error) {
	urls := []stm.URL{
		{"loc": "/users"},
		{"loc": "/users/sign-in"},
		{"loc": "/users/sign-up"},
		{"loc": "/users/confirm"},
		{"loc": "/users/forgot-password"},
		{"loc": "/users/unlock"},
	}
	return urls, nil
}

// Dashboard dashboard
func (p *Plugin) Dashboard(c *gin.Context) *web.Dropdown {
	if _, ok := c.Get(CurrentUser); !ok {
		return nil
	}
	return &web.Dropdown{
		Label: "auth.dashboard.title",
		Links: []*web.Link{
			&web.Link{Href: "/users/info", Label: "auth.users.info.title"},
			&web.Link{Href: "/users/change-password", Label: "auth.users.change-password.title"},
			nil,
			&web.Link{Href: "/users/logs", Label: "auth.users.logs.title"},
			nil,
			&web.Link{Href: "/attachments", Label: "auth.attachments.index.title"},
		},
	}
}

func init() {
	web.Register(&Plugin{})
}
