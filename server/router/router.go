package router

import (
	"github.com/codeyhj/auth-admin/server/handle"
	"github.com/labstack/echo/v4"
)

func AddRouter(e *echo.Echo) {
	g := e.Group("/admin")
	g.GET("/client/list", handle.HandleClient)
	g.GET("/token/list", handle.HandleToken)
	g.GET("/account/list", handle.HandleAccount)
	g.POST("/client/add", handle.HandleClientAdd)
	g.POST("/client/edit", handle.HandleClientEdit)
}
