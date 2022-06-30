package router

import (
	"github.com/codeyhj/auth-admin/server/handle"
	"github.com/labstack/echo/v4"
)

func AddRouter(e *echo.Echo) {

	g := e.Group("/api")

	g.GET("/token/list", handle.HandleToken)

	g.GET("/account/list", handle.HandleAccount)
	g.POST("/account/add", handle.HandleAccountAdd)
	g.POST("/account/edit", handle.HandleAccountEdit)

	g.GET("/client/list", handle.HandleClient)
	g.POST("/client/add", handle.HandleClientAdd)
	g.POST("/client/edit", handle.HandleClientEdit)

	e.Static("/assets", "dist/assets")

	e.File("*", "dist/index.html")

}
