package reading

import (
	"fmt"

	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine engine
type Engine struct {
	Db   *gorm.DB  `inject:""`
	I18n *web.I18n `inject:""`
	Jwt  *auth.Jwt `inject:""`
}

// RegisterWorker register worker
func (p *Engine) RegisterWorker() {

}

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	var books []Book
	if err := p.Db.Select([]string{"id"}).Find(&books).Error; err != nil {
		return nil, err
	}
	urls := []stm.URL{
		{"loc": "/reading/books"},
	}
	for _, b := range books {
		urls = append(
			urls,
			stm.URL{"loc": fmt.Sprintf("/reading/books/%d", b.ID)},
		)
	}
	return urls, nil
}

// Dashboard dashboard
func (p *Engine) Dashboard(c *gin.Context) *web.Dropdown {
	if _, ok := c.Get(auth.CurrentUser); !ok {
		return nil
	}
	dd := web.Dropdown{
		Label: "reading.dashboard.title",
		Links: []*web.Link{
			&web.Link{Href: "/reading/notes/my", Label: "reading.notes.my.title"},
		},
	}
	if admin, ok := c.Get(auth.IsAdmin); ok && admin.(bool) {
		dd.Links = append(
			dd.Links,
			&web.Link{Href: "/reading/admin/books", Label: "reading.admin.books.index.title"},
			&web.Link{Href: "/reading/admin/status", Label: "reading.admin.status.title"},
		)
	}
	return &dd
}

func init() {
	web.Register(&Engine{})
}
