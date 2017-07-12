package vpn

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/h2o/plugins/auth"
	"github.com/kapmahc/h2o/web"
)

// Mount web mount-points
func (p *Plugin) Mount(rt *gin.Engine) {
	ag := rt.Group("/ops/vpn", p.Jwt.MustAdminMiddleware)
	ag.GET("/users", auth.HTML(p.indexUsers))
	ag.GET("/users/new", auth.HTML(p.createUser))
	ag.POST("/users/new", auth.HTML(p.createUser))
	ag.GET("/users/edit/:id", auth.HTML(p.updateUser))
	ag.POST("/users/edit/:id", auth.HTML(p.updateUser))
	ag.GET("/users/reset-password/:id", auth.HTML(p.resetUserPassword))
	ag.POST("/users/reset-password/:id", auth.HTML(p.resetUserPassword))
	ag.DELETE("/users/:id", web.JSON(p.destroyUser))

	ag.GET("/logs", auth.HTML(p.indexLogs))

	ag.GET("/readme", auth.HTML(p.getReadme))

	rt.GET("/ops/vpn/users/change-password", auth.HTML(p.changeUserPassword))
	rt.POST("/ops/vpn/users/change-password", auth.HTML(p.changeUserPassword))

	api := rt.Group("/ops/vpn/api", p.tokenMiddleware)
	api.POST("/auth", web.JSON(p.apiAuth))
	api.POST("/connect", web.JSON(p.apiConnect))
	api.POST("/disconnect", web.JSON(p.apiDisconnect))
}
