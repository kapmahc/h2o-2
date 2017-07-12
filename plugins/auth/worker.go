package auth

import (
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/SermoDigital/jose/jws"
	"github.com/kapmahc/h2o/web"
	log "github.com/sirupsen/logrus"
	gomail "gopkg.in/gomail.v2"
)

const (
	actConfirm       = "confirm"
	actUnlock        = "unlock"
	actResetPassword = "reset-password"

	sendEmailJob = "auth.send-email"
)

// RegisterWorker register worker
func (p *Plugin) RegisterWorker() {
	p.Server.RegisterTask(sendEmailJob, p.doSendEmail)
}

func (p *Plugin) sendEmail(lng string, user *User, act string) {
	cm := jws.Claims{}
	cm.Set("act", act)
	cm.Set("uid", user.UID)
	tkn, err := p.Jwt.Sum(cm, time.Hour*6)
	if err != nil {
		log.Error(err)
		return
	}

	obj := struct {
		Home  string
		Token string
	}{
		Home:  web.Home(),
		Token: string(tkn),
	}
	subject, err := p.I18n.F(lng, fmt.Sprintf("auth.emails.%s.subject", act), obj)
	if err != nil {
		log.Error(err)
		return
	}
	body, err := p.I18n.F(lng, fmt.Sprintf("auth.emails.%s.body", act), obj)
	if err != nil {
		log.Error(err)
		return
	}

	// -----------------------
	task := tasks.Signature{
		Name: sendEmailJob,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: user.Email,
			},
			{
				Type:  "string",
				Value: subject,
			},
			{
				Type:  "string",
				Value: body,
			},
		},
	}

	if _, err := p.Server.SendTask(&task); err != nil {
		log.Error(err)
	}
}

func (p *Plugin) parseToken(lng, tkn, act string) (*User, error) {
	cm, err := p.Jwt.Validate([]byte(tkn))
	if err != nil {
		return nil, err
	}
	if act != cm.Get("act").(string) {
		return nil, p.I18n.E(lng, "errors.bad-action")
	}
	return p.Dao.GetUserByUID(cm.Get("uid").(string))
}

func (p *Plugin) doSendEmail(to, subject, body string) (interface{}, error) {
	if !(web.IsProduction()) {
		log.Debugf("send to %s: %s\n%s", to, subject, body)
		return "done", nil
	}
	smtp := make(map[string]interface{})
	if err := p.Settings.Get("site.smtp", &smtp); err != nil {
		return nil, err
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

	if err := dia.DialAndSend(msg); err != nil {
		return nil, err
	}

	return "done", nil
}
