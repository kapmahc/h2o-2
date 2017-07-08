package nut

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
	"golang.org/x/text/language"
)

// Controller base
type Controller struct {
	beego.Controller

	Locale string
}

// Prepare prepare
func (p *Controller) Prepare() {
	p.detectLocale()
	beego.ReadFromRequest(&p.Controller)
	p.Data["xsrf_token"] = p.XSRFToken()
	p.Data["xsrf"] = template.HTML(p.XSRFFormHTML())
}

// Abort abort http
func (p *Controller) Abort(code int, err error) {
	if err == nil {
		p.CustomAbort(code, strconv.Itoa(code))
	} else {
		p.CustomAbort(code, err.Error())
	}
}

// Bind bind params to form and valid it
func (p *Controller) Bind(fm interface{}) error {
	if err := p.ParseForm(fm); err != nil {
		return err
	}
	var va validation.Validation
	ok, err := va.Valid(fm)
	var msg []string
	if err != nil {
		return err
	}
	if !ok {
		for _, err := range va.Errors {
			msg = append(msg, fmt.Sprintf("%s: %s", err.Field, err.Message))
		}
		return errors.New(strings.Join(msg, "</li><li>"))
	}
	return nil
}

// Flash check error
func (p *Controller) Flash(err error, sgo, ego string) {
	if err == nil {
		p.Redirect(sgo, http.StatusFound)
		return
	}

	flash := beego.NewFlash()
	flash.Error(err.Error())
	flash.Store(&p.Controller)
	p.Redirect(ego, http.StatusFound)
}

// SetApplicationLayout using application layout
func (p *Controller) SetApplicationLayout() {
	p.Layout = "layouts/application.html"
}

// SetDashboardLayout using dashboard layout
func (p *Controller) SetDashboardLayout() {
	p.Layout = "layouts/dashboard.html"
}

func (p *Controller) detectLocale() {
	const key = "locale"
	write := false

	// 1. Check URL arguments.
	lang := p.Input().Get(key)

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = p.Ctx.GetCookie(key)
	} else {
		write = true
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := p.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			lang = al[:5] // Only compare first 5 letters.
		}
		write = true
	}

	// 4. Default language is English.
	tag, err := language.Parse(lang)
	if err != nil {
		beego.Error(err)
	}
	lang = tag.String()
	if !i18n.IsExist(lang) {
		lang = language.AmericanEnglish.String()
		write = true
	}

	// Save language information in cookies.
	if write {
		p.Ctx.SetCookie(key, lang, 1<<31-1, "/")
	}

	// Set language properties.
	p.Locale = lang
	p.Data[key] = lang
}

func init() {
	beego.AddFuncMap("t", T)
	beego.AddFuncMap("dict", func(args ...interface{}) (map[string]interface{}, error) {
		size := len(args)
		if size%2 != 0 {
			return nil, errors.New("invalid dict call")
		}
		dict := make(map[string]interface{}, size/2)
		for i := 0; i < size; i += 2 {
			key, ok := args[i].(string)
			if !ok {
				return nil, errors.New("dict keys must be strings")
			}
			dict[key] = args[i+1]
		}
		return dict, nil
	})
}
