package nut

import "github.com/astaxie/beego"

// Controller base
type Controller struct {
	beego.Controller

	Locale string
}

// Prepare prepare
func (p *Controller) Prepare() {
	p.detectLocale()
}
