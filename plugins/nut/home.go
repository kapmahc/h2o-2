package nut

// HomeController home
type HomeController struct {
	Controller
}

// Index home
// @router / [get]
func (p *HomeController) Index() {
	p.TplName = "nut/index.html"
}
