package route

import (
	"github.com/labstack/echo"
	"test/Handler"
)

func Router(e *echo.Echo) {

	push := e.Group("/push")
	//"user" URI
	push.GET("",Handler.UserListHandler)
	push.POST("",Handler.AddUser)
	push.GET("/:id",Handler.UserHandler)
	return
}
