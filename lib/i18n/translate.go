package i18n

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
)

// T translate message
func T(lang, code string, args ...interface{}) string {
	msg, ok := get(lang, code)
	if !ok {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}

// E translate error
func E(lang, code string, args ...interface{}) error {
	msg, ok := get(lang, code)
	if !ok {
		return errors.New(msg)
	}
	return fmt.Errorf(msg, args...)
}

// F translate template
func F(lang, code string, arg interface{}) (string, error) {
	msg, ok := get(lang, code)
	if !ok {
		return msg, nil
	}
	tpl, err := template.New("").Parse(msg)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, arg)
	return buf.String(), err
}
