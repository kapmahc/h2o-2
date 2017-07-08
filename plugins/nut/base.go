package nut

import (
	"github.com/astaxie/beego"
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
}

// DetectLocale detect locale from http request
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
