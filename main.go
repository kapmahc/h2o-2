package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/kapmahc/h2o/routers"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	beego.Run()
}
