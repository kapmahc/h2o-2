package main

import (
	"path"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego/toolbox"
	"github.com/kapmahc/h2o/plugins/nut"
	_ "github.com/kapmahc/h2o/routers"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+path.Join("tmp", "h2o.log")+`", "maxdays":180, "perm":"0600"}`)
	if err := nut.CheckDb(); err != nil {
		beego.Error(err)
		return
	}
	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()
}
