package auth

import (
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
	"github.com/kapmahc/h2o/plugins/nut"
)

const (
	emailFormTpl = "auth/users/email-form.html"
)

// UsersController users
type UsersController struct {
	Controller
}

type fmSignIn struct {
	Email    string `form:"email" valid:"Email;Required"`
	Password string `form:"password" valid:"Required"`
}

// SignIn sign in
// @router /sign-in [get,post]
func (p *UsersController) SignIn() {
	if p.Ctx.Request.Method == http.MethodPost {
		var fm fmSignIn
		var user *User
		ip := p.Ctx.Input.IP()

		err := p.Bind(&fm)
		if err == nil {
			user, err = GetUserByEmail(fm.Email)
		}
		if err == nil {
			if !nut.Chk([]byte(fm.Password), []byte(user.Password)) {
				AddLog(user.ID, ip, p.Locale, "auth.logs.sign-in-failed")
				err = nut.E(p.Locale, "auth.errors.user-bad-password")
			}
		}
		if err == nil {
			if !user.IsConfirm() {
				err = nut.E(p.Locale, "auth.errors.user-is-not-confirm")
			}
		}
		if err == nil {
			if user.IsLock() {
				err = nut.E(p.Locale, "auth.errors.user-is-lock")
			}
		}

		if err == nil {
			AddLog(user.ID, ip, p.Locale, "auth.logs.sign-in-success")
			p.SetSession("uid", user.UID)
			p.Success(nut.T(p.Locale, "auth.messages.sign-in-success"), "/")
		} else {
			p.Fail(err, p.signInPath())
		}
		return
	}
	// http get
	p.SetApplicationLayout()
	p.Data["title"] = i18n.Tr(p.Locale, "auth.users.sign-in.title")
	p.TplName = "auth/users/sign-in.html"
}

type fmSignUp struct {
	Name                 string `form:"name" valid:"Required;MaxSize(32)"`
	Email                string `form:"email" valid:"Email;MaxSize(255)"`
	Password             string `form:"password" valid:"MinSize(6)"`
	PasswordConfirmation string `form:"passwordConfirmation" `
}

func (p *fmSignUp) Valid(v *validation.Validation) {
	if p.Password != p.PasswordConfirmation {
		v.SetError("PasswordConfirmation", "Passwords not match")
	}
}

// SignUp sign up
// @router /sign-up [get,post]
func (p *UsersController) SignUp() {
	if p.Ctx.Request.Method == http.MethodPost {
		var fm fmSignUp
		var user *User
		ip := p.Ctx.Input.IP()

		err := p.Bind(&fm)
		if err == nil {
			user, err = AddEmailUser(fm.Email, fm.Password, ip, p.Locale)
		}
		if err == nil {
			err = p.sendEmail(user, p.Locale, actConfirm)
		}
		if err == nil {
			p.Success(nut.T(p.Locale, "auth.messages.email-for-confirm"), p.signInPath())
		} else {
			p.Fail(err, p.URLFor("auth.UsersController.SignUp"))
		}
		return
	}
	// http get
	p.SetApplicationLayout()
	p.Data["title"] = i18n.Tr(p.Locale, "auth.users.sign-up.title")
	p.TplName = "auth/users/sign-up.html"
}

func (p *UsersController) signInPath() string {
	return p.URLFor("auth.UsersController.SignIn")
}

type fmEmail struct {
	Email string `form:"email" valid:"Email;MaxSize(255)"`
}

// Confirm confirm
// @router /confirm [get,post]
func (p *UsersController) Confirm() {
	if p.Ctx.Request.Method == http.MethodPost {
		var fm fmEmail
		var user *User

		err := p.Bind(&fm)
		if err == nil {
			user, err = GetUserByEmail(fm.Email)
		}
		if err == nil {
			if user.IsConfirm() {
				err = nut.E(p.Locale, "auth.errors.user-already-confirm")
			}
		}
		if err == nil {
			err = p.sendEmail(user, p.Locale, actConfirm)
		}

		if err == nil {
			p.Success(nut.T(p.Locale, "auth.messages.email-for-confirm"), p.signInPath())
		} else {
			p.Fail(err, p.URLFor("auth.UsersController.Confirm"))
		}
		return
	}
	// http get
	p.SetApplicationLayout()
	p.Data["title"] = i18n.Tr(p.Locale, "auth.users.confirm.title")
	p.TplName = emailFormTpl
}

// GetConfirmToken confirm token
// @router /confirm/:token [get]
func (p *UsersController) GetConfirmToken() {
	user, err := p.parseToken(p.Locale, p.Ctx.Input.Param(":token"), actConfirm)
	if err == nil {
		if user.IsConfirm() {
			err = nut.E(p.Locale, "auth.errors.user-already-confirm")
		}
	}
	if err == nil {
		err = ConfirmUser(user.ID, p.Ctx.Input.IP(), p.Locale)
	}
	if err == nil {
		p.Success(nut.T(p.Locale, "auth.messages.confirm-success"), p.signInPath())
	} else {
		p.Fail(err, p.signInPath())
	}
}

// Unlock unlock
// @router /unlock [get,post]
func (p *UsersController) Unlock() {
	if p.Ctx.Request.Method == http.MethodPost {
		var fm fmEmail
		var user *User

		err := p.Bind(&fm)
		if err == nil {
			user, err = GetUserByEmail(fm.Email)
		}
		if err == nil {
			if !user.IsLock() {
				err = nut.E(p.Locale, "auth.errors.user-is-not-lock")
			}
		}
		if err == nil {
			err = p.sendEmail(user, p.Locale, actUnlock)
		}

		if err == nil {
			p.Success(nut.T(p.Locale, "auth.messages.email-for-unlock"), p.signInPath())
		} else {
			p.Fail(err, p.URLFor("auth.UsersController.Unlock"))
		}
		return
	}
	// http get
	p.SetApplicationLayout()
	p.Data["title"] = i18n.Tr(p.Locale, "auth.users.unlock.title")
	p.TplName = emailFormTpl
}

// GetUnlockToken unlock token
// @router /unlock/:token [get]
func (p *UsersController) GetUnlockToken() {
	user, err := p.parseToken(p.Locale, p.Ctx.Input.Param(":token"), actUnlock)
	if err == nil {
		if !user.IsLock() {
			err = nut.E(p.Locale, "auth.errors.user-is-not-lock")
		}
	}
	if err == nil {
		o := orm.NewOrm()
		_, err = o.QueryTable(new(User)).Filter("id", user.ID).Update(orm.Params{
			"locked_at": nil,
		})
	}
	if err == nil {
		err = AddLog(user.ID, p.Ctx.Input.IP(), p.Locale, "auth.logs.unlock")
	}
	if err == nil {
		p.Success(nut.T(p.Locale, "auth.messages.unlock-success"), p.signInPath())
	} else {
		p.Fail(err, p.signInPath())
	}
}

// ForgotPassword forgot password
// @router /forogot-password [get,post]
func (p *UsersController) ForgotPassword() {
	if p.Ctx.Request.Method == http.MethodPost {
		var fm fmEmail
		var user *User

		err := p.Bind(&fm)
		if err == nil {
			user, err = GetUserByEmail(fm.Email)
		}
		if err == nil {
			err = p.sendEmail(user, p.Locale, actResetPassword)
		}
		if err == nil {
			p.Success(nut.T(p.Locale, "auth.messages.email-for-reset-password"), p.signInPath())
		} else {
			p.Fail(err, p.URLFor("auth.UsersController.ForgotPassword"))
		}
		return
	}
	// http get
	p.SetApplicationLayout()
	p.Data["title"] = i18n.Tr(p.Locale, "auth.users.forgot-password.title")
	p.TplName = emailFormTpl
}

type fmResetPassword struct {
	Password             string `form:"password" valid:"MinSize(6)"`
	PasswordConfirmation string `form:"passwordConfirmation" `
}

func (p *fmResetPassword) Valid(v *validation.Validation) {
	if p.Password != p.PasswordConfirmation {
		v.SetError("PasswordConfirmation", "Passwords not match")
	}
}

// ResetPassword reset password
// @router /reset-password/:token [get,post]
func (p *UsersController) ResetPassword() {
	if p.Ctx.Request.Method == http.MethodPost {
		token := p.Ctx.Input.Param(":token")
		var fm fmResetPassword
		var user *User

		err := p.Bind(&fm)
		if err == nil {
			user, err = p.parseToken(p.Locale, token, actResetPassword)
		}

		if err == nil {
			o := orm.NewOrm()
			_, err = o.QueryTable(new(User)).Filter("id", user.ID).Update(orm.Params{
				"password": string(nut.Sum([]byte(fm.Password))),
			})
		}
		if err == nil {
			err = AddLog(user.ID, p.Ctx.Input.IP(), p.Locale, "auth.logs.reset-password")
		}

		if err == nil {
			p.Success(nut.T(p.Locale, "auth.messages.reset-password-success"), p.signInPath())
		} else {
			p.Fail(err, p.URLFor("auth.UsersController.ResetPassword", ":token", token))
		}
		return
	}
	// http get
	p.SetApplicationLayout()
	p.Data["title"] = i18n.Tr(p.Locale, "auth.users.reset-password.title")
	p.TplName = "auth/users/reset-password.html"
}
