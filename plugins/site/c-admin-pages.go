package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/h2o/web"
)

var (
	cardLocs = []interface{}{"carousel", "circle", "square"}
	defLogo  = "data:image/gif;base64,R0lGODlhAQABAIAAAHd3dwAAACH5BAAAAAAALAAAAAABAAEAAAICRAEAOw=="
)

func (p *Plugin) indexAdminCards(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.cards.index.title")
	tpl := "site-admin-cards-index"
	var cards []web.Card
	if err := p.Db.Order("loc ASC, sort_order DESC").Find(&cards).Error; err != nil {
		return tpl, err
	}
	data["items"] = cards
	return tpl, nil
}

type fmCard struct {
	Title     string `form:"title" binding:"required,max=255"`
	Href      string `form:"href" binding:"required,max=255"`
	Loc       string `form:"loc" binding:"required,max=16"`
	Action    string `form:"action" binding:"required,max=32"`
	Logo      string `form:"logo" binding:"required,max=255"`
	Summary   string `form:"summary" binding:"required,max=2048"`
	SortOrder int    `form:"sortOrder"`
}

func (p *Plugin) createAdminCard(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	data["locs"] = cardLocs
	data["sortOrders"] = sortOrders
	tpl := "site-admin-cards-new"
	if c.Request.Method == http.MethodPost {
		var fm fmCard
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&web.Card{
			Loc:       fm.Loc,
			Title:     fm.Title,
			Summary:   fm.Summary,
			Logo:      fm.Logo,
			Href:      fm.Href,
			Action:    fm.Action,
			SortOrder: fm.SortOrder,
		}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/admin/cards")
		return "", nil
	}
	return tpl, nil
}

func (p *Plugin) updateAdminCard(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	data["locs"] = cardLocs
	data["sortOrders"] = sortOrders
	tpl := "site-admin-cards-edit"
	id := c.Param("id")

	var item web.Card
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item
	if c.Request.Method == http.MethodPost {
		var fm fmCard
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&web.Card{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"loc":        fm.Loc,
				"href":       fm.Href,
				"sort_order": fm.SortOrder,
				"title":      fm.Title,
				"summary":    fm.Summary,
				"action":     fm.Action,
				"logo":       fm.Logo,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/admin/cards")
		return "", nil
	}

	return tpl, nil
}

func (p *Plugin) destroyAdminCard(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(web.Card{}).Error
	return gin.H{}, err
}
