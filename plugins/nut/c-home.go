package nut

// HomeController home
type HomeController struct {
	Controller
}

// GetInstall install
// @router /install [get]
func (p *HomeController) GetInstall() {
	p.TplName = "nut/install.html"
}

// PostInstall install
// @router /install [post]
func (p *HomeController) PostInstall() {

}
