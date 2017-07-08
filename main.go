package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/kapmahc/h2o/routers"
)

func main() {
	beego.Run()
}
