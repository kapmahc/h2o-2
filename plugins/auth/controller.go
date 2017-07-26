package auth

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/h2o/plugins/nut"
)

// Controller controller
type Controller struct {
	nut.Controller
	CurrentUser *User
	IsAdmin     bool
}

// Prepare prepare
func (p *Controller) Prepare() {
	p.Controller.Prepare()
	p.setCurrentUser()
}

func (p *Controller) setCurrentUser() {
	uid := p.GetSession("uid")
	if uid == nil {
		return
	}

	user, err := GetUserByUID(uid.(string))
	if err != nil {
		beego.Error(err)
		return
	}
	isAdmin := Is(user.ID, RoleAdmin)

	p.CurrentUser = user
	p.IsAdmin = isAdmin
	p.Data["currentUser"] = map[string]interface{}{
		"name":    user.Name,
		"isAdmin": isAdmin,
	}
}
