package nut

import (
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
)

type fmInstall struct {
	Title                string `form:"title" valid:"Required"`
	SubTitle             string `form:"subTitle" valid:"Required"`
	Name                 string `form:"name" valid:"MaxSize(32)"`
	Email                string `form:"email" valid:"Email;MaxSize(255)"`
	Password             string `form:"password" valid:"Required"`
	PasswordConfirmation string `form:"passwordConfirmation" `
}

func (p *fmInstall) Valid(v *validation.Validation) {
	if p.Password != p.PasswordConfirmation {
		v.SetError("PasswordConfirmation", "Passwords not match")
	}
}

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
		var fm fmInstall
		err := p.Bind(&fm)
		if err == nil {
			err = SetMessage(p.Locale, "site.title", fm.Title)
		}
		if err == nil {
			err = SetMessage(p.Locale, "site.sub-title", fm.SubTitle)
		}

		var user *User
		ip := p.Ctx.Input.IP()
		if err == nil {
			_, err = AddEmailUser(fm.Email, fm.Password, ip, p.Locale)
		}
		if err == nil {
			err = ConfirmUser(user.ID, ip, p.Locale)
		}

		p.Flash(err)
		if err == nil {
			p.Redirect("/", http.StatusFound)
		}
	}
	p.SetApplicationLayout()

	p.Data["title"] = i18n.Tr(p.Locale, "site.install.title")
	p.TplName = "nut/install.html"
}
