package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Redirect redirect
func Redirect(f func(*gin.Context) (u string, e error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, e := f(c)
		if e != nil {
			c.String(http.StatusInternalServerError, e.Error())
			return
		}
		c.Redirect(http.StatusFound, u)
	}
}

// JSON json render
func JSON(f func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, e := f(c)
		if e != nil {
			c.String(http.StatusInternalServerError, e.Error())
			return
		}
		c.JSON(http.StatusOK, v)
	}
}
