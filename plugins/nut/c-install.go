package nut

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
)

// Install install
// @router /install [get,post]
func (p *HomeController) Install() {
	o := orm.NewOrm()
	count, err := o.QueryTable(new(User)).Count()
	if err != nil {
		p.Abort(http.StatusInternalServerError, err)
	}
	if count > 0 {
		p.Abort(http.StatusForbidden, nil)
	}

	if p.Ctx.Request.Method == http.MethodPost {

	}
	p.SetApplicationLayout()

	p.Data["title"] = i18n.Tr(p.Locale, "site.install.title")
	beego.Debug(p.Locale, p.Data, i18n.ListLangs())
	p.TplName = "nut/install.html"
}
