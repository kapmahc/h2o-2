package auth

import (
	"fmt"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/astaxie/beego"
	"github.com/kapmahc/h2o/plugins/nut"
	gomail "gopkg.in/gomail.v2"
)

const (
	actConfirm       = "confirm"
	actUnlock        = "unlock"
	actResetPassword = "reset-password"
)

func (p *UsersController) sendEmail(user *User, lang, act string) error {
	cm := jws.Claims{}
	cm.Set("act", act)
	cm.Set("uid", user.UID)
	tkn, err := nut.JwtSum(cm, time.Hour*6)
	if err != nil {
		return err
	}

	obj := struct {
		Home  string
		Token string
	}{
		Home:  nut.Home(),
		Token: string(tkn),
	}
	subject, err := nut.H(lang, fmt.Sprintf("auth.emails.%s.subject", act), obj)
	if err != nil {
		return err
	}
	body, err := nut.H(lang, fmt.Sprintf("auth.emails.%s.body", act), obj)
	if err != nil {
		return err
	}

	// -----------------------

	go func() {
		if err := p.doSendEmail(user.Email, subject, body); err != nil {
			beego.Error(err)
		}
	}()
	return nil
}

func (p *UsersController) parseToken(lng, tkn, act string) (*User, error) {
	cm, err := nut.JwtValidate([]byte(tkn))
	if err != nil {
		return nil, err
	}
	if act != cm.Get("act").(string) {
		return nil, nut.E(lng, "errors.bad-action")
	}
	return GetUserByUID(cm.Get("uid").(string))
}

func (p *UsersController) doSendEmail(to, subject, body string) error {
	if beego.BConfig.RunMode != beego.PROD {
		beego.Debug("send to:", to, subject, body)
		return nil
	}
	smtp := make(map[string]interface{})
	if err := nut.Get("site.smtp", &smtp); err != nil {
		return err
	}

	sender := smtp["username"].(string)
	msg := gomail.NewMessage()
	msg.SetHeader("From", sender)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dia := gomail.NewDialer(
		smtp["host"].(string),
		smtp["port"].(int),
		sender,
		smtp["password"].(string),
	)

	return dia.DialAndSend(msg)
}
