package routers

import (
	"github.com/kapmahc/h2o/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
