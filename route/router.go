package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"test/Handler"
)

func Router(e *echo.Echo) {
	//Setting middleware func
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut,http.MethodPost,http.MethodDelete},
	}))
	push := e.Group("/push")
	//"user" URI
	push.GET("",Handler.UserListHandler)
	push.POST("",Handler.AddUser)
	push.GET("/:id",Handler.UserHandler)

	return
}
