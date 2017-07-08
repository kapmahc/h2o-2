package nut

// Install install
// @router /install [get]
func (p *HomeController) Install() {
	p.TplName = "nut/install.html"
}
