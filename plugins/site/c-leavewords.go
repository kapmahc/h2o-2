package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/h2o/web"
)

func (p *Plugin) indexLeaveWords(c *gin.Context, lang string, data gin.H) (string, error) {
	var items []LeaveWord
	err := p.Db.Order("created_at DESC").Find(&items).Error
	data["title"] = p.I18n.T(lang, "site.leave-words.index.title")
	data["items"] = items
	return "site-leave-words-index", err
}

type fmLeaveWord struct {
	Body string `form:"body" binding:"required,max=800"`
	Type string `form:"type" binding:"required,max=8"`
}

func (p *Plugin) createLeaveWord(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.leave-words.index.title")
	tpl := "site-leave-words-new"
	if c.Request.Method == http.MethodPost {
		var fm fmLeaveWord
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		if err := p.Db.Create(&LeaveWord{Type: fm.Type, Body: fm.Body}).Error; err != nil {
			return tpl, err
		}
		data[web.NOTICE] = p.I18n.T(lang, "success")
	}
	return tpl, nil
}

func (p *Plugin) destroyLeaveWord(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(LeaveWord{}).Error
	return gin.H{}, err
}
