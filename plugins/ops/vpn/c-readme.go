package vpn

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/h2o/plugins/auth"
	"github.com/kapmahc/h2o/web"
	"github.com/spf13/viper"
)

func (p *Plugin) getReadme(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "ops.vpn.readme.title")
	tpl := "ops-vpn-readme"
	data["user"] = c.MustGet(auth.CurrentUser)
	data["name"] = viper.Get("server.name")
	data["home"] = web.Home()
	data["port"] = 1194
	data["network"] = "10.18.0"

	token, err := p.generateToken(10)
	if err != nil {
		return tpl, err
	}
	data["token"] = string(token)
	return tpl, nil
}
