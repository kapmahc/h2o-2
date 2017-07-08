package nut

import "github.com/astaxie/beego"

// UsersController users controller
type UsersController struct {
	beego.Controller
}

// GetSignIn sign in
// @router /sign-in [get]
func (p *UsersController) GetSignIn() {

}
