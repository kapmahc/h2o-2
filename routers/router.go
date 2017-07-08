package routers

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/h2o/plugins/auth"
	"github.com/kapmahc/h2o/plugins/erp"
	"github.com/kapmahc/h2o/plugins/forum"
	"github.com/kapmahc/h2o/plugins/mall"
	"github.com/kapmahc/h2o/plugins/ops/mail"
	"github.com/kapmahc/h2o/plugins/ops/vpn"
	"github.com/kapmahc/h2o/plugins/pos"
	"github.com/kapmahc/h2o/plugins/reading"
	"github.com/kapmahc/h2o/plugins/site"
	"github.com/kapmahc/h2o/plugins/survey"
)

func init() {
	beego.Include(&site.HomeController{})

	for k, v := range map[string]beego.ControllerInterface{
		"/users":       &auth.UsersController{},
		"/attachments": &auth.AttachmentsController{},
		"/admin":       &site.AdminController{},

		"/forum":    &forum.Controller{},
		"/reading":  &reading.Controller{},
		"/survey":   &survey.Controller{},
		"/ops/vpn":  &vpn.Controller{},
		"/ops/mail": &mail.Controller{},
		"/erp":      &erp.Controller{},
		"/mall":     &mall.Controller{},
		"/pos":      &pos.Controller{},
	} {
		beego.AddNamespace(beego.NewNamespace(k, beego.NSInclude(v)))
	}

}
