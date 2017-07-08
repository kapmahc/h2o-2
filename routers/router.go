package routers

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/h2o/plugins/erp"
	"github.com/kapmahc/h2o/plugins/forum"
	"github.com/kapmahc/h2o/plugins/mall"
	"github.com/kapmahc/h2o/plugins/nut"
	"github.com/kapmahc/h2o/plugins/ops/mail"
	"github.com/kapmahc/h2o/plugins/ops/vpn"
	"github.com/kapmahc/h2o/plugins/pos"
	"github.com/kapmahc/h2o/plugins/reading"
	"github.com/kapmahc/h2o/plugins/survey"
)

func init() {
	for _, v := range []beego.ControllerInterface{
		&nut.UsersController{},
		&nut.HomeController{},
	} {
		beego.Include(v)
	}

	for k, v := range map[string]beego.ControllerInterface{
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
