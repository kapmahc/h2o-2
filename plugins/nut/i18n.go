package nut

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

// T translate string
func T(lang, code string, args ...interface{}) string {
	msg, err := GetMessage(lang, code)
	if err != nil {
		return i18n.Tr(lang, code, args...)
	}
	return fmt.Sprintf(msg, args...)
}

// E translate error
func E(lang, code string, args ...interface{}) error {
	msg, err := GetMessage(lang, code)
	if err != nil {
		return errors.New(i18n.Tr(lang, code, args...))
	}
	return fmt.Errorf(msg, args...)
}

// H translate html template
func H(lang, code string, obj interface{}) (string, error) {
	msg, err := GetMessage(lang, code)
	if err != nil {
		msg = i18n.Tr(lang, code)
	}
	tpl, err := template.New("").Parse(msg)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, obj)
	return buf.String(), err
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
