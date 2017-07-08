package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego/toolbox"
	"github.com/kapmahc/h2o/plugins/base"
	_ "github.com/kapmahc/h2o/routers"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	if err := base.CheckDb(); err != nil {
		beego.Error(err)
		return
	}
	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()
}
