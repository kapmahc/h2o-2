package site

import "github.com/gin-gonic/gin"

func (p *Plugin) getDashboard(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "header.dashboard")
	return "site-dashboard", nil
}

func (p *Plugin) getHome(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "header.home")
	return "site-home", nil
}
