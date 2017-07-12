package mail

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/h2o/plugins/auth"
	"github.com/kapmahc/h2o/web"
)

// Mount web mount-points
func (p *Plugin) Mount(rt *gin.Engine) {
	ag := rt.Group("/ops/mail", p.Jwt.MustAdminMiddleware)

	ag.GET("/users", auth.HTML(p.indexUsers))
	ag.GET("/users/new", auth.HTML(p.createUser))
	ag.POST("/users/new", auth.HTML(p.createUser))
	ag.GET("/users/edit/:id", auth.HTML(p.updateUser))
	ag.POST("/users/edit/:id", auth.HTML(p.updateUser))
	ag.GET("/users/reset-password/:id", auth.HTML(p.resetUserPassword))
	ag.POST("/users/reset-password/:id", auth.HTML(p.resetUserPassword))
	ag.DELETE("/users/:id", web.JSON(p.destroyUser))

	rt.GET("/ops/mail/users/change-password", auth.HTML(p.changeUserPassword))
	rt.POST("/ops/mail/users/change-password", auth.HTML(p.changeUserPassword))

	ag.GET("/domains", auth.HTML(p.indexDomains))
	ag.GET("/domains/new", auth.HTML(p.createDomain))
	ag.POST("/domains/new", auth.HTML(p.createDomain))
	ag.GET("/domains/edit/:id", auth.HTML(p.updateDomain))
	ag.POST("/domains/edit/:id", auth.HTML(p.updateDomain))
	ag.DELETE("/domains/:id", web.JSON(p.destroyDomain))

	ag.GET("/aliases", auth.HTML(p.indexAliases))
	ag.GET("/aliases/new", auth.HTML(p.createAlias))
	ag.POST("/aliases/new", auth.HTML(p.createAlias))
	ag.GET("/aliases/edit/:id", auth.HTML(p.updateAlias))
	ag.POST("/aliases/edit/:id", auth.HTML(p.updateAlias))
	ag.DELETE("/aliases/:id", web.JSON(p.destroyAlias))

	ag.GET("/readme", auth.HTML(p.getReadme))
}
