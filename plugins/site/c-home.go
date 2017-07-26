package site

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
	"github.com/kapmahc/h2o/plugins/auth"
	"github.com/kapmahc/h2o/plugins/nut"
)

// HomeController home
type HomeController struct {
	auth.Controller
}

// Index home
// @router / [get]
func (p *HomeController) Index() {
	p.SetApplicationLayout()
	beego.Debug(p.Data)
	p.TplName = "site/index.html"
}

// Search search
// @router /search [post]
func (p *HomeController) Search() {
	p.SetApplicationLayout()
	p.TplName = "site/search.html"
}

// Dashboard dashboard
// @router /dashboard [get]
func (p *HomeController) Dashboard() {
	p.SetDashboardLayout()
	p.TplName = "site/dashboard.html"
}

type fmInstall struct {
	Title                string `form:"title" valid:"Required"`
	SubTitle             string `form:"subTitle" valid:"Required"`
	Name                 string `form:"name" valid:"Required;MaxSize(32)"`
	Email                string `form:"email" valid:"Email;MaxSize(255)"`
	Password             string `form:"password" valid:"MinSize(6)"`
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
	// check database
	o := orm.NewOrm()
	count, err := o.QueryTable(new(auth.User)).Count()
	if err != nil {
		p.Abort(http.StatusInternalServerError, err)
	}
	if count > 0 {
		p.Abort(http.StatusForbidden, nil)
	}
	// http post
	if p.Ctx.Request.Method == http.MethodPost {
		var fm fmInstall
		err := p.Bind(&fm)
		if err == nil {
			err = nut.SetMessage(p.Locale, "site.title", fm.Title)
		}
		if err == nil {
			err = nut.SetMessage(p.Locale, "site.sub-title", fm.SubTitle)
		}

		var user *auth.User
		ip := p.Ctx.Input.IP()
		if err == nil {
			user, err = auth.AddEmailUser(fm.Name, fm.Email, fm.Password, ip, p.Locale)
		}
		if err == nil {
			err = auth.ConfirmUser(user.ID, ip, p.Locale)
		}

		if err == nil {
			for _, role := range []string{"root", "admin", "member"} {
				if err = auth.Allow(user.ID, role, auth.DefaultResourceType, auth.DefaultResourceID, 10, 0, 0); err != nil {
					break
				}
			}
		}

		if err == nil {
			p.Success(nut.T(p.Locale, "auth.messages.confirm-success"), p.URLFor("auth.UsersController.SignIn"))
			p.Redirect(p.URLFor("auth.UsersController.SignIn"), http.StatusFound)
		} else {
			p.Fail(err, p.URLFor("site.HomeController.Install"))
		}
		return
	}

	// http get
	p.SetApplicationLayout()
	p.Data["title"] = i18n.Tr(p.Locale, "site.install.title")
	p.TplName = "site/install.html"
}
