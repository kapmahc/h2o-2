package auth

import "github.com/kapmahc/h2o/plugins/nut"

// UsersController users
type UsersController struct {
	nut.Controller
}

// SignIn sign in
// @router /sign-in [get,post]
func (p *UsersController) SignIn() {

}
