package mail

import "github.com/gin-gonic/gin"

func (p *Plugin) getReadme(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "ops.mail.readme.title")
	tpl := "ops-mail-readme"

	return tpl, nil
}
