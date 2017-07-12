package main

import (
	"log"

	_ "github.com/kapmahc/h2o/plugins/erp"
	_ "github.com/kapmahc/h2o/plugins/forum"
	_ "github.com/kapmahc/h2o/plugins/mall"
	_ "github.com/kapmahc/h2o/plugins/ops/mail"
	_ "github.com/kapmahc/h2o/plugins/ops/vpn"
	_ "github.com/kapmahc/h2o/plugins/pos"
	_ "github.com/kapmahc/h2o/plugins/reading"
	_ "github.com/kapmahc/h2o/plugins/site"
	_ "github.com/kapmahc/h2o/plugins/suvery"
	"github.com/kapmahc/h2o/web"
	_ "github.com/lib/pq"
)

func main() {
	if err := web.Main(); err != nil {
		log.Fatal(err)
	}
}
