package mail

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) getReadme(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "ops.mail.readme.title")
	tpl := "ops-mail-readme"

	return tpl, nil
}
