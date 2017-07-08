package nut

import (
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"golang.org/x/text/language"
)

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

func init() {
	const ext = ".ini"
	if err := filepath.Walk("locales", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if !info.IsDir() && filepath.Ext(name) == ext {
			lang := name[:len(name)-len(ext)]
			beego.Debug("find locale", lang)
			return i18n.SetMessage(lang, path)
		}
		return nil
	}); err != nil {
		beego.Error(err)
	}
}
