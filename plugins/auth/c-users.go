package auth

import "github.com/gin-gonic/gin"

func (p *Plugin) indexUsers(c *gin.Context, l string, d gin.H) (string, error) {
	var users []User
	d["title"] = p.I18n.T(l, "auth.users.index.title")
	err := p.Db.
		Select([]string{"name", "logo", "home"}).
		Order("last_sign_in_at DESC").
		Find(&users).Error
	d["users"] = users
	return "auth-users-index", err
}
