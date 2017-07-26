package auth

import (
	"github.com/kapmahc/axe"
	"github.com/kapmahc/axe/base"
)

// Mount web mount points
func (p *Plugin) Mount(rt *axe.Router) {
	rt.GET("/users/sign-in", p.Wrapper.HTML(base.LayoutApplication, "auth/users/sign-in", p.getSignIn))
}
